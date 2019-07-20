package pack

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	. "satellite/utils"
	"strings"
	"sync"
)

func PackBase64One(srcfile string) (r string, err error) {
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
	// third, split the data slice
	ss, err := SplitByte(data, ConstBase64BufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	// fourth, we can call Base64Encrypt function
	wg := &sync.WaitGroup{}
	rr := make([]string, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go Base64EncrypteGo(string(v), &rr[k], wg)
	}
	wg.Wait()
	dest := strings.Join(rr, "")
	// fifth, fill the packet struct
	_, srcname := filepath.Split(srcfile)
	if len([]byte(srcname)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	head := TPackBase64One{}
	head.Name = make([]byte, 32)
	head.Size = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(srcname))
	BytesCopy(&(head.Size), IntToBytes(len(dest)))
	// finally, return result
	var s []string
	s = append(s, string(head.Name))
	s = append(s, string(head.Size))
	s = append(s, dest)
	r = strings.Join(s, "")
	return r, err
}

func Base64EncrypteGo(str string, r *string, wg *sync.WaitGroup) {
	*r = Base64Encrypt(str)
	wg.Done()
}

func Base64Encrypt(str string) string {
	s := []byte(str)
	r := base64.StdEncoding.EncodeToString(s)
	return r
}
