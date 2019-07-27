package unpack

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	. "satellite/utils"
	"strings"
	"sync"
)

func UnpackBase64(srcfile string, destpath string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// first, open the file
	file, err := os.Open(srcfile)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	defer file.Close()
	// second, read file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	_, srcname := filepath.Split(srcfile)
	// third, new one header
	h := TUnpackBase64{}
	h.Name = make([]byte, 32)
	h.Author = make([]byte, 16)
	h.Type = make([]byte, 8)
	h.Number = make([]byte, 4)
	// fourth, read the header
	rd := bytes.NewReader(data)
	_, err = rd.Read(h.Name)
	if err != nil {
		log.Println("Error read header name:", err)
		return err
	}
	s := make([]byte, 32)
	BytesCopy(&s, []byte(srcname))
	if !bytes.Equal(h.Name, s) {
		log.Println("Error read header name:", err)
		return err
	}
	_, err = rd.Read(h.Author)
	if err != nil {
		log.Println("Error read header author:", err)
		return err
	}
	s = make([]byte, 16)
	BytesCopy(&s, []byte("Alopex6414"))
	if !bytes.Equal(h.Author, s) {
		log.Println("Error read header author:", err)
		return err
	}
	_, err = rd.Read(h.Type)
	if err != nil {
		log.Println("Error read header type:", err)
		return err
	}
	s = make([]byte, 8)
	BytesCopy(&s, []byte("BASE64"))
	if !bytes.Equal(h.Type, s) {
		log.Println("Error read header type:", err)
		return err
	}
	_, err = rd.Read(h.Number)
	if err != nil {
		log.Println("Error read header number:", err)
		return err
	}
	size := BytesToInt(h.Number)
	// fifth, read every one file in packet
	for i := 0; i < size; i++ {
		// six, read the header
		hh := TUnpackBase64One{}
		hh.Name = make([]byte, 32)
		hh.Size = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return err
		}
		_, err = rd.Read(hh.Size)
		if err != nil {
			log.Println("Error read header size:", err)
			return err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return err
		}
		// eight, run unpack one file
		wg.Add(1)
		go UnpackBase64OneGo(s, hh, destpath, wg)
	}
	wg.Wait()
	return err
}

func UnpackBase64OneGo(data []byte, head TUnpackBase64One, destpath string, wg *sync.WaitGroup) (err error) {
	err = UnpackBase64One(data, head, destpath)
	if err != nil {
		log.Println("Error Base64 Unpack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func UnpackBase64One(data []byte, head TUnpackBase64One, destpath string) (err error) {
	// initial, fill the name
	var s []byte
	for _, v := range head.Name {
		if v == byte(0) {
			break
		}
		s = append(s, v)
	}
	destfile := destpath + string(s)
	// first, split the data slice
	ss, err := SplitByte(data, ConstBase64BufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	size := len(data) % ConstBase64BufferSize
	if size != 0 {
		last := len(data) / ConstBase64BufferSize
		ss[last] = append(ss[last][:0], ss[last][:size]...)
	}
	// second, we can call Base64Decrypt function
	wg := &sync.WaitGroup{}
	rr := make([]string, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go Base64DecryptGo(string(v), &rr[k], wg)
	}
	wg.Wait()
	dest := strings.Join(rr, "")
	// fourth, create the origin file
	err = ioutil.WriteFile(destfile, []byte(dest), 0644)
	if err != nil {
		log.Println("Error write to destfile:", err)
		return err
	}
	return err
}

func Base64DecryptGo(str string, r *string, wg *sync.WaitGroup) {
	*r = Base64Decrypt(str)
	wg.Done()
}

func Base64Decrypt(str string) string {
	s, _ := base64.StdEncoding.DecodeString(str)
	r := string(s)
	return r
}
