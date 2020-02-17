package pack

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/global"
	. "satellite/utils"
	"strings"
	"sync"
	"sync/atomic"
)

// PackBase64 function
// input source file list and dest package path, output error information
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// dest file name suffix can be any type such as '.pak', '.dat', even none is ok
// return err indicate the success or failure function execute
func PackBase64(src []string, dest string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// clear global variable
	atomic.StoreInt64(&Done, 0)
	// first, split the pre-crypt files
	r := make([]string, len(src)+4)
	for k, v := range src {
		wg.Add(1)
		go PackBase64OneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, check goroutine whether success or not
	for i := 0; i < len(src); i++ {
		if r[i+4] == "" {
			s := fmt.Sprintf("Error base64 pack one file: %v", src[i])
			err = errors.New(s)
			return err
		}
	}
	// third, fill the header
	head := TPackBase64{}
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
	BytesCopy(&(head.Type), []byte("BASE64"))
	BytesCopy(&(head.Number), IntToBytes(len(src)))
	r[0] = string(head.Name)
	r[1] = string(head.Author)
	r[2] = string(head.Type)
	r[3] = string(head.Number)
	// finally, write to dest file
	s := strings.Join(r, "")
	err = ioutil.WriteFile(dest, []byte(s), 0644)
	if err != nil {
		log.Println("Error write base64 file:", err)
	}
	return err
}

// PackBase64WorkCalculate function
// it will calculate the total work value which you input files
// it will be call in progress pack files
// input src files same as you pack src files
// output work value and err
// return err indicate the success or failure function execute
func PackBase64WorkCalculate(src []string) (work int64, err error) {
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
		if size%Base64BufferSize != 0 {
			padding := Base64BufferSize - size%Base64BufferSize
			size += padding
		}
		sum += size
	}
	work = sum
	return work, err
}

// PackBase64OneGo function
// input source file, return value pointer and wait group pointer
// it will pack one file through goroutine
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// r should input string pointer and it will fill in return value
// wg is a flag to control different goroutine sync
// return err indicate the success or failure function execute
func PackBase64OneGo(src string, r *string, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*r, err = PackBase64One(src)
	if err != nil {
		log.Println("Error base64 pack one file:", err)
		return err
	}
	return err
}

// PackBase64One function
// this function is the base function of PackBase64OneGo
func PackBase64One(src string) (r string, err error) {
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
	// third, split the data slice
	ss, err := SplitByte(data, Base64BufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	size := len(data) % Base64BufferSize
	if size != 0 {
		last := len(data) / Base64BufferSize
		ss[last] = append(ss[last][:0], ss[last][:size]...)
	}
	// fourth, we can call Base64Encrypt function
	wg := &sync.WaitGroup{}
	rr := make([]string, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go Base64EncryptGo(string(v), &rr[k], wg)
	}
	wg.Wait()
	dest := strings.Join(rr, "")
	// fifth, fill the packet struct
	_, name := filepath.Split(src)
	if len([]byte(name)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	head := TPackBase64One{}
	head.Name = make([]byte, 32)
	head.Size = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(name))
	BytesCopy(&(head.Size), IntToBytes(len(dest)))
	// finally, return result
	var s []string
	s = append(s, string(head.Name))
	s = append(s, string(head.Size))
	s = append(s, dest)
	r = strings.Join(s, "")
	return r, err
}

// Base64EncryptGo function
// input source file, return value pointer and wait group pointer
// it will encrypt one file through goroutine
// src file must be string, you can open file and read it through io
// dest should input string pointer and it will fill in return value after encrypt
// wg is a flag to control different goroutine sync
// return err indicate the success or failure function execute
func Base64EncryptGo(str string, r *string, wg *sync.WaitGroup) {
	defer wg.Done()
	*r = Base64Encrypt(str)
	atomic.AddInt64(&Done, 1)
}

// Base64Encrypt function
// it common with function Base64EncryptGo
func Base64Encrypt(str string) string {
	s := []byte(str)
	r := base64.StdEncoding.EncodeToString(s)
	return r
}
