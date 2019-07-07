package pack

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	. "satellite/utils"
	"sync"
	"time"
)

func PackAES(srcfilelist []string, destfile string) (err error) {
	wg := &sync.WaitGroup{}
	r := make([][]byte, len(srcfilelist))
	for k, v := range srcfilelist {
		wg.Add(1)
		go PackAESOneGo(v, &r[k], wg)
	}
	wg.Wait()
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
	r = []byte{}
	err = nil
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
	if len([]byte(srcfile)) > 32 {
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
	head.Name = []byte(srcfile)
	head.Key = key
	head.OriginSize = IntToBytes(len(data))
	head.CryptSize = IntToBytes(len(dest))
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
	dest = []byte{}
	err = nil
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
	padding := size - len(src) % size
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, text...)
}

func SplitByte(data []byte, size int) (r [][]byte, err error) {
	rd := bytes.NewReader(data)
	for {
		s := make([]byte, size)
		switch n, err := rd.Read(s); true {
		case n < 0:
			log.Println("Error read byte:", err)
			return r, err
		case n == 0:
			return r, nil
		case n > 0:
			r = append(r, s)
		}
	}
}