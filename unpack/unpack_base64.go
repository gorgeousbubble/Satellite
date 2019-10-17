package unpack

import (
	"bytes"
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
	"sync/atomic"
)

func UnpackBase64(src string, dest string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// clear global variable
	atomic.StoreInt64(&Done, 0)
	// first, open the file
	file, err := os.Open(src)
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
	_, name := filepath.Split(src)
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
	BytesCopy(&s, []byte(name))
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
		go UnpackBase64OneGo(s, hh, dest, wg)
	}
	wg.Wait()
	return err
}

func UnpackBase64ToFile(src string, target string, dest string) (err error) {
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// first, open the file
	file, err := os.Open(src)
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
	_, name := filepath.Split(src)
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
	BytesCopy(&s, []byte(name))
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
		// eight, when it is target file, then run unpack one file
		if target == string(bytes.Trim(hh.Name, "\x00")) {
			err = UnpackBase64One(s, hh, dest)
			if err != nil {
				log.Println("Error unpack base64 one to file:", err)
				return err
			}
			return nil
		}
	}
	return err
}

func UnpackBase64ToMemory(src string, target string, dest *[]byte) (err error) {
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// first, open the file
	file, err := os.Open(src)
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
	_, name := filepath.Split(src)
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
	BytesCopy(&s, []byte(name))
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
		// eight, when it is target file, then run unpack one file
		if target == string(bytes.Trim(hh.Name, "\x00")) {
			var r string
			err = UnpackBase64OneToMemory(s, &r)
			if err != nil {
				log.Println("Error unpack base64 one to memroy:", err)
				return err
			}
			*dest = []byte(r)
			return nil
		}
	}
	return err
}

func UnpackBase64ExtractInfo(src string, dest *[]string, sz *[]int) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
	_, name := filepath.Split(src)
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
	BytesCopy(&s, []byte(name))
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
		// eight, extract packet information
		*dest = append(*dest, string(bytes.Trim(hh.Name, "\x00")))
		*sz = append(*sz, BytesToInt(hh.Size))
	}
	return err
}

func UnpackBase64WorkCalculate(src string) (work int64, err error) {
	var sum int64
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return work, err
	}
	defer file.Close()
	// second, read file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file:", err)
		return work, err
	}
	_, name := filepath.Split(src)
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
		return work, err
	}
	s := make([]byte, 32)
	BytesCopy(&s, []byte(name))
	if !bytes.Equal(h.Name, s) {
		log.Println("Error read header name:", err)
		return work, err
	}
	_, err = rd.Read(h.Author)
	if err != nil {
		log.Println("Error read header author:", err)
		return work, err
	}
	s = make([]byte, 16)
	BytesCopy(&s, []byte("Alopex6414"))
	if !bytes.Equal(h.Author, s) {
		log.Println("Error read header author:", err)
		return work, err
	}
	_, err = rd.Read(h.Type)
	if err != nil {
		log.Println("Error read header type:", err)
		return work, err
	}
	s = make([]byte, 8)
	BytesCopy(&s, []byte("BASE64"))
	if !bytes.Equal(h.Type, s) {
		log.Println("Error read header type:", err)
		return work, err
	}
	_, err = rd.Read(h.Number)
	if err != nil {
		log.Println("Error read header number:", err)
		return work, err
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
			return work, err
		}
		_, err = rd.Read(hh.Size)
		if err != nil {
			log.Println("Error read header size:", err)
			return work, err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return work, err
		}
		// eight, calculate file size sum
		sum += int64(BytesToInt(hh.Size))
	}
	work = sum
	return work, err
}

func UnpackBase64OneToMemory(data []byte, dest *string) (err error) {
	// first, split the data slice
	ss, err := SplitByte(data, Base64BufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	size := len(data) % Base64BufferSize
	if size != 0 {
		last := len(data) / Base64BufferSize
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
	*dest = strings.Join(rr, "")
	return err
}

func UnpackBase64OneGo(data []byte, head TUnpackBase64One, dest string, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	err = UnpackBase64One(data, head, dest)
	if err != nil {
		log.Println("Error base64 unpack one file:", err)
		return err
	}
	return err
}

func UnpackBase64One(data []byte, head TUnpackBase64One, path string) (err error) {
	// initial, fill the name
	var s []byte
	for _, v := range head.Name {
		if v == byte(0) {
			break
		}
		s = append(s, v)
	}
	file := path + string(s)
	// first, split the data slice
	ss, err := SplitByte(data, Base64BufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	size := len(data) % Base64BufferSize
	if size != 0 {
		last := len(data) / Base64BufferSize
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
	err = ioutil.WriteFile(file, []byte(dest), 0644)
	if err != nil {
		log.Println("Error write to dest file:", err)
		return err
	}
	return err
}

func Base64DecryptGo(str string, r *string, wg *sync.WaitGroup) {
	defer wg.Done()
	*r = Base64Decrypt(str)
	atomic.AddInt64(&Done, 1)
}

func Base64Decrypt(str string) string {
	s, _ := base64.StdEncoding.DecodeString(str)
	r := string(s)
	return r
}
