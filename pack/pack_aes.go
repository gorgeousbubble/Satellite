package pack

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/global"
	. "satellite/utils"
	"sync"
	"time"
)

func PackAES(srcfilelist []string, destfile string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// first, split the pre-crypt files
	r := make([][]byte, len(srcfilelist)+4)
	for k, v := range srcfilelist {
		wg.Add(1)
		go PackAESOneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, check goroutine whether success or not
	for i := 0; i < len(srcfilelist); i++ {
		if bytes.Equal(r[i+4], []byte("")) {
			err = errors.New("error aes packet one file")
			return err
		}
	}
	// third, fill the header
	head := TPackAES{}
	head.Name = make([]byte, 32)
	head.Author = make([]byte, 16)
	head.Type = make([]byte, 8)
	head.Number = make([]byte, 4)
	_, destname := filepath.Split(destfile)
	if len([]byte(destname)) > 32 {
		log.Println("Error dest file name length:", err)
		return
	}
	BytesCopy(&(head.Name), []byte(destname))
	BytesCopy(&(head.Author), []byte("Alopex6414"))
	BytesCopy(&(head.Type), []byte("AES"))
	BytesCopy(&(head.Number), IntToBytes(len(srcfilelist)))
	r[0] = head.Name
	r[1] = head.Author
	r[2] = head.Type
	r[3] = head.Number
	// finally, write to dest file
	s := bytes.Join(r, []byte(""))
	err = ioutil.WriteFile(destfile, s, 0644)
	if err != nil {
		log.Println("Error Write AES:", err)
	}
	return err
}

func PackAESOneGo(srcfile string, r *[]byte, wg *sync.WaitGroup) (err error) {
	*r, err = PackAESOne(srcfile)
	if err != nil {
		log.Println("Error AES Pack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func PackAESOne(srcfile string) (r []byte, err error) {
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
	key := make([]byte, 16)
	_, err = rand.Read(key)
	if err != nil {
		log.Println("Error generate random key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, ConstAESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	// fifth, we can call AESEncrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go AESEncryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// sixth, fill the packet struct
	_, srcname := filepath.Split(srcfile)
	if len([]byte(srcname)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	if len(key) > 16 {
		log.Println("Error key length:", err)
		return
	}
	head := TPackAESOne{}
	head.Name = make([]byte, 32)
	head.Key = make([]byte, 16)
	head.OriginSize = make([]byte, 4)
	head.CryptSize = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(srcname))
	BytesCopy(&(head.Key), key)
	BytesCopy(&(head.OriginSize), IntToBytes(len(data)))
	BytesCopy(&(head.CryptSize), IntToBytes(len(dest)))
	/*// fourth, we can call AESEncrypt function
	dest, err := AESEncrypt(data, key)
	if err != nil {
		log.Println("Error AES Encrypt data:", err)
		return r, err
	}*/
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

func AESEncryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = AESEncrypt(src, key)
	if err != nil {
		log.Println("Error AES Encrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func AESEncrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 16, 24, 32
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error key length:", err)
		return dest, err
	}
	// calculate block size
	blockSize := block.BlockSize()
	// fill block data
	src = PKCS7Padding(src, blockSize)
	// encrypt mode
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// create slice
	dest = make([]byte, len(src))
	// encrypt
	blockMode.CryptBlocks(dest, src)
	return dest, err
}

func PKCS7Padding(src []byte, size int) []byte {
	var padding int
	if len(src)%size != 0 {
		padding = size - len(src)%size
	}
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, text...)
}
