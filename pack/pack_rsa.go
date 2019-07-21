package pack

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/utils"
	"sync"
)

func PackRSA(srcfilelist []string, destfile string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// first, split the pre-crypt files
	r := make([][]byte, len(srcfilelist)+3)
	for k, v := range srcfilelist {
		wg.Add(1)
		go PackRSAOneGo(v, &r[k+3], wg)
	}
	wg.Wait()
	// second, fill the header
	head := TPackRSA{}
	head.Name = make([]byte, 32)
	head.Author = make([]byte, 16)
	head.Number = make([]byte, 4)
	_, destname := filepath.Split(destfile)
	if len([]byte(destname)) > 32 {
		log.Println("Error dest file name length:", err)
		return
	}
	BytesCopy(&(head.Name), []byte(destname))
	BytesCopy(&(head.Author), []byte("Alopex6414"))
	BytesCopy(&(head.Number), IntToBytes(len(srcfilelist)))
	r[0] = head.Name
	r[1] = head.Author
	r[2] = head.Number
	// third, write to dest file
	s := bytes.Join(r, []byte(""))
	err = ioutil.WriteFile(destfile, s, 0644)
	if err != nil {
		log.Println("Error Write RSA:", err)
	}
	return err
}

func PackRSAOneGo(srcfile string, r *[]byte, wg *sync.WaitGroup) (err error) {
	*r, err = PackRSAOne(srcfile)
	if err != nil {
		log.Println("Error RSA Pack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func PackRSAOne(srcfile string) (r []byte, err error) {
	// first, open the file
	file, err := os.Open(srcfile)
	if err != nil {
		log.Println("Error open file:", err)
		return r, err
	}
	defer file.Close()
	// second, read file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file:", err)
		return r, err
	}
	// third, generate rsa key
	var pri []byte
	var pub []byte
	err = GenRSAKey2Memory(&pri, &pub, 1024)
	if err != nil {
		log.Println("Error generate rsa key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, ConstRSABufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	// fifth, we can call RSAEncrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go RSAEncryptGo(v, pub, &rr[k], wg)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// sixth, fill the packet struct
	_, srcname := filepath.Split(srcfile)
	if len([]byte(srcname)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	head := TPackRSAOne{}
	head.Name = make([]byte, 32)
	head.Key = make([]byte, 1024)
	head.OriginSize = make([]byte, 4)
	head.CryptSize = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(srcname))
	BytesCopy(&(head.Key), pri)
	BytesCopy(&(head.OriginSize), IntToBytes(len(data)))
	BytesCopy(&(head.CryptSize), IntToBytes(len(dest)))
	// finally, return result
	var s [][]byte
	s = append(s, head.Name)
	s = append(s, head.Key)
	s = append(s, head.OriginSize)
	s = append(s, head.CryptSize)
	s = append(s, dest)
	r = bytes.Join(s, []byte(""))
	return r, err
}

func RSAEncryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = RSAEncrypt(src, key)
	if err != nil {
		log.Println("Error RSA Encrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func RSAEncrypt(src, key []byte) (dest []byte, err error) {
	block, _ := pem.Decode(key)
	if block == nil {
		err = errors.New("RSA Public Key Error")
		return dest, err
	}
	pi, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return dest, err
	}
	pub := pi.(*rsa.PublicKey)
	dest, err = rsa.EncryptPKCS1v15(rand.Reader, pub, src)
	if err != nil {
		return dest, err
	}
	return dest, err
}