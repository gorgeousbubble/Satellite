package pack

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/utils"
	"sync"
	"time"
)

func PackDES(srcfilelist []string, destfile string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// first, split the pre-crypt files
	r := make([][]byte, len(srcfilelist)+3)
	for k, v := range srcfilelist {
		wg.Add(1)
		go PackDESOneGo(v, &r[k+3], wg)
	}
	wg.Wait()
	// second, fill the header
	head := TPackDES{}
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
		log.Println("Error Write DES:", err)
	}
	return err
}

func Pack3DESOneGo(srcfile string, r *[]byte, wg *sync.WaitGroup) (err error) {
	*r, err = Pack3DESOne(srcfile)
	if err != nil {
		log.Println("Error 3DES Pack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func Pack3DESOne(srcfile string) (r []byte, err error) {
	rand.Seed(time.Now().UnixNano())
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
	// third, generate random key
	key := make([]byte, 24)
	_, err = rand.Read(key)
	if err != nil {
		log.Println("Error generate random key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, ConstDESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	// fifth, we can call TripleDESEncryptGo function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go TripleDESEncryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// sixth, fill the packet struct
	_, srcname := filepath.Split(srcfile)
	if len([]byte(srcname)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	if len(key) > 24 {
		log.Println("Error key length:", err)
		return
	}
	head := TPack3DESOne{}
	head.Name = make([]byte, 32)
	head.Key = make([]byte, 24)
	head.OriginSize = make([]byte, 4)
	head.CryptSize = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(srcname))
	BytesCopy(&(head.Key), key)
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

func PackDESOneGo(srcfile string, r *[]byte, wg *sync.WaitGroup) (err error) {
	*r, err = PackDESOne(srcfile)
	if err != nil {
		log.Println("Error DES Pack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func PackDESOne(srcfile string) (r []byte, err error) {
	rand.Seed(time.Now().UnixNano())
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
	// third, generate random key
	key := make([]byte, 8)
	_, err = rand.Read(key)
	if err != nil {
		log.Println("Error generate random key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, ConstDESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	// fifth, we can call DESEncrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go DESEncryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// sixth, fill the packet struct
	_, srcname := filepath.Split(srcfile)
	if len([]byte(srcname)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	if len(key) > 8 {
		log.Println("Error key length:", err)
		return
	}
	head := TPackDESOne{}
	head.Name = make([]byte, 32)
	head.Key = make([]byte, 8)
	head.OriginSize = make([]byte, 4)
	head.CryptSize = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(srcname))
	BytesCopy(&(head.Key), key)
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

func TripleDESEncryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = TripleDESEncrypt(src, key)
	if err != nil {
		log.Println("Error 3DES Encrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func TripleDESEncrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 24
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		log.Println("Error key length:", err)
		return dest, err
	}
	// calculate block size
	blockSize := block.BlockSize()
	// fill block data
	src = PKCS5Padding(src, blockSize)
	// encrypt mode
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// create slice
	dest = make([]byte, len(src))
	// encrypt
	blockMode.CryptBlocks(dest, src)
	return dest, err
}

func DESEncryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = DESEncrypt(src, key)
	if err != nil {
		log.Println("Error DES Encrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func DESEncrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 8
	block, err := des.NewCipher(key)
	if err != nil {
		log.Println("Error key length:", err)
		return dest, err
	}
	// calculate block size
	blockSize := block.BlockSize()
	// fill block data
	src = PKCS5Padding(src, blockSize)
	// encrypt mode
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// create slice
	dest = make([]byte, len(src))
	// encrypt
	blockMode.CryptBlocks(dest, src)
	return dest, err
}

func PKCS5Padding(src []byte, size int) []byte {
	var padding int
	if len(src)%size != 0 {
		padding = size - len(src) % size
	}
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, text...)
}
