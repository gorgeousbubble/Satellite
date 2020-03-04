package pack

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/global"
	. "satellite/utils"
	"sync"
	"sync/atomic"
)

// PackRSA function
// input source file list and dest package path, output error information
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// dest file name suffix can be any type such as '.pak', '.dat', even none is ok
// return err indicate the success or failure function execute
func PackRSA(src []string, dest string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// clear global variable
	atomic.StoreInt64(&Done, 0)
	// first, split the pre-crypt files
	r := make([][]byte, len(src)+4)
	for k, v := range src {
		wg.Add(1)
		go PackRSAOneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, check goroutine whether success or not
	for i := 0; i < len(src); i++ {
		if bytes.Equal(r[i+4], []byte("")) {
			s := fmt.Sprintf("Error rsa pack one file: %v", src[i])
			err = errors.New(s)
			return err
		}
	}
	// third, fill the header
	head := TPackRSA{}
	head.Name = make([]byte, 32)
	head.Author = make([]byte, 16)
	head.Type = make([]byte, 8)
	head.Number = make([]byte, 4)
	_, name := filepath.Split(dest)
	if len([]byte(name)) > 32 {
		log.Println("Error dest file name length:", err)
		return
	}
	BytesCopy(&(head.Name), []byte(name))
	BytesCopy(&(head.Author), []byte("Alopex6414"))
	BytesCopy(&(head.Type), []byte("RSA"))
	BytesCopy(&(head.Number), IntToBytes(len(src)))
	r[0] = head.Name
	r[1] = head.Author
	r[2] = head.Type
	r[3] = head.Number
	// finally, write to dest file
	s := bytes.Join(r, []byte(""))
	err = ioutil.WriteFile(dest, s, 0644)
	if err != nil {
		log.Println("Error write rsa file:", err)
	}
	return err
}

// PackRSAWorkCalculate function
// it will calculate the total work value which you input files
// it will be call in progress pack files
// input src files same as you pack src files
// output work value and err
// return err indicate the success or failure function execute
func PackRSAWorkCalculate(src []string) (work int64, err error) {
	var sum int64
	if len(src) == 0 {
		err = errors.New("Pack file list is empty.")
		return work, err
	}
	for _, v := range src {
		var size int64
		err = filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			size = info.Size()
			return err
		})
		if err != nil {
			log.Println("Error calculate work:", err)
			return work, err
		}
		if size%RSAPacketSize != 0 {
			padding := RSAPacketSize - size%RSAPacketSize
			size += padding
		}
		sum += size
	}
	work = sum
	return work, err
}

// PackRSAOneGo function
// input source file, return value pointer and wait group pointer
// it will pack one file through goroutine
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// r should input byte slice pointer and it will fill in return value
// wg is a flag to control different goroutine sync
// return err indicate the success or failure function execute
func PackRSAOneGo(src string, r *[]byte, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*r, err = PackRSAOne(src)
	if err != nil {
		log.Println("Error rsa pack one file:", err)
		return err
	}
	return err
}

// PackRSAOne function
// it the base function of PackRSAOne
func PackRSAOne(src string) (r []byte, err error) {
	// first, open the file
	file, err := os.Open(src)
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
	ss, err := SplitByte(data, RSAPacketSize)
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
	_, name := filepath.Split(src)
	if len([]byte(name)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	head := TPackRSAOne{}
	head.Name = make([]byte, 32)
	head.Key = make([]byte, 1024)
	head.OriginSize = make([]byte, 4)
	head.CryptSize = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(name))
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

// RSAEncryptGo function
// input source file, encrypt key, return value pointer and wait group pointer
// it will encrypt one file through goroutine
// src file must be byte slice, you can open file and read it through io
// key is a 128bit number which used by rsa, here is 16 bit byte slice
// dest should input byte slice pointer and it will fill in return value after encrypt
// wg is a flag to control different goroutine sync
// return err indicate the success or failure function execute
func RSAEncryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*dest, err = RSAEncrypt(src, key)
	if err != nil {
		log.Println("Error rsa encrypt data:", err)
		return err
	}
	atomic.AddInt64(&Done, 1)
	return err
}

// RSAEncrypt function
// it the base function of RSAEncryptGo
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
