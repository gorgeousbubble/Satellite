package pack

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/global"
	. "satellite/utils"
	"strings"
	"sync"
)

func PackBase64(src []string, dest string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// first, split the pre-crypt files
	r := make([]string, len(src)+4)
	for k, v := range src {
		wg.Add(1)
		go PackBase64OneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, fill the header
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
	// third, write to dest file
	s := strings.Join(r, "")
	err = ioutil.WriteFile(dest, []byte(s), 0644)
	if err != nil {
		log.Println("Error write base64 file:", err)
	}
	return err
}

func PackBase64OneGo(src string, r *string, wg *sync.WaitGroup) (err error) {
	*r, err = PackBase64One(src)
	if err != nil {
		log.Println("Error base64 pack one file:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

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
	ss, err := SplitByte(data, ConstBase64BufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	size := len(data) % ConstBase64BufferSize
	if size != 0 {
		last := len(data) / ConstBase64BufferSize
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

func Base64EncryptGo(str string, r *string, wg *sync.WaitGroup) {
	*r = Base64Encrypt(str)
	wg.Done()
}

func Base64Encrypt(str string) string {
	s := []byte(str)
	r := base64.StdEncoding.EncodeToString(s)
	return r
}
