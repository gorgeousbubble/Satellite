package pack

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/global"
	. "satellite/utils"
	"sync"
	"sync/atomic"
	"time"
)

// Pack3DES function
// input source file list and dest package path, output error information
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// dest file name suffix can be any type such as '.pak', '.dat', even none is ok
// return err indicate the success or failure function execute
func Pack3DES(src []string, dest string) (err error) {
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
		go Pack3DESOneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, check goroutine whether success or not
	for i := 0; i < len(src); i++ {
		if bytes.Equal(r[i+4], []byte("")) {
			s := fmt.Sprintf("Error 3des pack one file: %v", src[i])
			err = errors.New(s)
			return err
		}
	}
	// third, fill the header
	head := TPack3DES{}
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
	BytesCopy(&(head.Type), []byte("3DES"))
	BytesCopy(&(head.Number), IntToBytes(len(src)))
	r[0] = head.Name
	r[1] = head.Author
	r[2] = head.Type
	r[3] = head.Number
	// finally, write to dest file
	s := bytes.Join(r, []byte(""))
	err = ioutil.WriteFile(dest, s, 0644)
	if err != nil {
		log.Println("Error write 3des file:", err)
	}
	return err
}

func PackDES(src []string, dest string) (err error) {
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
		go PackDESOneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, check goroutine whether success or not
	for i := 0; i < len(src); i++ {
		if bytes.Equal(r[i+4], []byte("")) {
			s := fmt.Sprintf("Error des pack one file: %v", src[i])
			err = errors.New(s)
			return err
		}
	}
	// third, fill the header
	head := TPackDES{}
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
	BytesCopy(&(head.Type), []byte("DES"))
	BytesCopy(&(head.Number), IntToBytes(len(src)))
	r[0] = head.Name
	r[1] = head.Author
	r[2] = head.Type
	r[3] = head.Number
	// finally, write to dest file
	s := bytes.Join(r, []byte(""))
	err = ioutil.WriteFile(dest, s, 0644)
	if err != nil {
		log.Println("Error write des file:", err)
	}
	return err
}

func PackDESWorkCalculate(src []string) (work int64, err error) {
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
		if size%DESBufferSize != 0 {
			padding := DESBufferSize - size%DESBufferSize
			size += padding
		}
		sum += size
	}
	work = sum
	return work, err
}

func Pack3DESOneGo(src string, r *[]byte, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*r, err = Pack3DESOne(src)
	if err != nil {
		log.Println("Error 3des pack one file:", err)
		return err
	}
	return err
}

func Pack3DESOne(src string) (r []byte, err error) {
	rand.Seed(time.Now().UnixNano())
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
	// third, generate random key
	key := make([]byte, 24)
	_, err = rand.Read(key)
	if err != nil {
		log.Println("Error generate random key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, DESBufferSize)
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
	_, name := filepath.Split(src)
	if len([]byte(name)) > 32 {
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
	BytesCopy(&(head.Name), []byte(name))
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

func PackDESOneGo(src string, r *[]byte, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*r, err = PackDESOne(src)
	if err != nil {
		log.Println("Error des pack one file:", err)
		return err
	}
	return err
}

func PackDESOne(src string) (r []byte, err error) {
	rand.Seed(time.Now().UnixNano())
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
	// third, generate random key
	key := make([]byte, 8)
	_, err = rand.Read(key)
	if err != nil {
		log.Println("Error generate random key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, DESBufferSize)
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
	_, name := filepath.Split(src)
	if len([]byte(name)) > 32 {
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
	BytesCopy(&(head.Name), []byte(name))
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
	defer wg.Done()
	*dest, err = TripleDESEncrypt(src, key)
	if err != nil {
		log.Println("Error 3des encrypt data:", err)
		return err
	}
	atomic.AddInt64(&Done, 1)
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
	defer wg.Done()
	*dest, err = DESEncrypt(src, key)
	if err != nil {
		log.Println("Error des encrypt data:", err)
		return err
	}
	atomic.AddInt64(&Done, 1)
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
		padding = size - len(src)%size
	}
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, text...)
}
