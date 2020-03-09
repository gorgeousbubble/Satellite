package unpack

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
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

// UnpackAES function
// This function mainly used for unpack aes package.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// return err indicate the success or failure function execute
func UnpackAES(src string, dest string) (err error) {
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
	h := TUnpackAES{}
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
	BytesCopy(&s, []byte("AES"))
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
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return err
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			log.Println("Error read header key:", err)
			return err
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			log.Println("Error read header origin size:", err)
			return err
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			log.Println("Error read header crypt size:", err)
			return err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return err
		}
		// eight, run unpack one file
		wg.Add(1)
		go UnpackAESOneGo(s, hh, dest, wg)
	}
	wg.Wait()
	return err
}

// UnpackAESConfine function
// This function is mainly used for unpack aes package with restrict go routine.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// return err indicate the success or failure function execute
func UnpackAESConfine(src string, dest string) (err error) {
	wg := &sync.WaitGroup{}
	ch := make(chan interface{}, ConfineFiles)
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
	h := TUnpackAES{}
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
	BytesCopy(&s, []byte("AES"))
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
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return err
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			log.Println("Error read header key:", err)
			return err
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			log.Println("Error read header origin size:", err)
			return err
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			log.Println("Error read header crypt size:", err)
			return err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return err
		}
		// eight, run unpack one file
		wg.Add(1)
		ch <- struct{}{}
		go UnpackAESOneConfineGo(s, hh, dest, wg, ch)
	}
	wg.Wait()
	return err
}

// UnpackAESToFile function
// This function is mainly used for unpack aes package to file.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// target string is the file which you want to decrypt from package. for instance, if the original name of file is 'capture.png',
// you should fill target segment with 'capture.png'
// return err indicate the success or failure function execute
func UnpackAESToFile(src string, target string, dest string) (err error) {
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
	h := TUnpackAES{}
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
	BytesCopy(&s, []byte("AES"))
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
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return err
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			log.Println("Error read header key:", err)
			return err
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			log.Println("Error read header origin size:", err)
			return err
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			log.Println("Error read header crypt size:", err)
			return err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return err
		}
		// eight, when it is target file, then run unpack one file
		if target == string(bytes.Trim(hh.Name, "\x00")) {
			err = UnpackAESOne(s, hh, dest)
			if err != nil {
				log.Println("Error unpack aes one to file:", err)
				return err
			}
			return nil
		}
	}
	return err
}

// UnpackAESToFileConfine function
// This function is mainly used for unpack aes package to file with restrict go routine.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// target string is the file which you want to decrypt from package. for instance, if the original name of file is 'capture.png',
// you should fill target segment with 'capture.png'
// return err indicate the success or failure function execute
func UnpackAESToFileConfine(src string, target string, dest string) (err error) {
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
	h := TUnpackAES{}
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
	BytesCopy(&s, []byte("AES"))
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
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return err
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			log.Println("Error read header key:", err)
			return err
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			log.Println("Error read header origin size:", err)
			return err
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			log.Println("Error read header crypt size:", err)
			return err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return err
		}
		// eight, when it is target file, then run unpack one file
		if target == string(bytes.Trim(hh.Name, "\x00")) {
			err = UnpackAESOneConfine(s, hh, dest)
			if err != nil {
				log.Println("Error unpack aes one to file:", err)
				return err
			}
			return nil
		}
	}
	return err
}

// UnpackAESToMemory function
// This function is mainly used for unpack aes package to memory.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest is a slice which used to receive decrypt data. You can send '[]byte' slice address here.
// target string is the file which you want to decrypt from package. for instance, if the original name of file is 'capture.png',
// you should fill target segment with 'capture.png'
// return err indicate the success or failure function execute
func UnpackAESToMemory(src string, target string, dest *[]byte) (err error) {
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
	h := TUnpackAES{}
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
	BytesCopy(&s, []byte("AES"))
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
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return err
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			log.Println("Error read header key:", err)
			return err
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			log.Println("Error read header origin size:", err)
			return err
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			log.Println("Error read header crypt size:", err)
			return err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return err
		}
		// eight, when it is target file, then run unpack one file
		if target == string(bytes.Trim(hh.Name, "\x00")) {
			err = UnpackAESOneToMemory(s, hh, dest)
			if err != nil {
				log.Println("Error unpack aes one to memroy:", err)
				return err
			}
			return nil
		}
	}
	return err
}

// UnpackAESExtractInfo function
// This function is mainly used for check verbose information of package.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest string slice will return the files name in package.
// sz int slice will return the file number in package.
// algorithm will return which algorithm used by encrypt package.
// return err indicate the success or failure function execute
func UnpackAESExtractInfo(src string, dest *[]string, sz *[]int) (err error) {
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
	h := TUnpackAES{}
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
	BytesCopy(&s, []byte("AES"))
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
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return err
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			log.Println("Error read header key:", err)
			return err
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			log.Println("Error read header origin size:", err)
			return err
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			log.Println("Error read header crypt size:", err)
			return err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return err
		}
		// eight, extract packet information
		*dest = append(*dest, string(bytes.Trim(hh.Name, "\x00")))
		*sz = append(*sz, BytesToInt(hh.OriginSize))
	}
	return err
}

// UnpackAESWorkCalculate function
// This function is mainly used for calculate the total work of unpack process.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// algorithm return the algorithm type which used in unpack
// work return the total work value of unpack process.
// return err indicate the success or failure function execute
func UnpackAESWorkCalculate(src string) (work int64, err error) {
	var sum int64
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return work, err
	}
	// second, read file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file:", err)
		return work, err
	}
	_, name := filepath.Split(src)
	// third, new one header
	h := TUnpackAES{}
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
	BytesCopy(&s, []byte("AES"))
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
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)
		_, err = rd.Read(hh.Name)
		if err != nil {
			log.Println("Error read header name:", err)
			return work, err
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			log.Println("Error read header key:", err)
			return work, err
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			log.Println("Error read header origin size:", err)
			return work, err
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			log.Println("Error read header crypt size:", err)
			return work, err
		}
		// seven, read the body
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			log.Println("Error read body:", err)
			return work, err
		}
		// eight, calculate file size sum
		sum += int64(BytesToInt(hh.CryptSize))
	}
	work = sum
	return work, err
}

// UnpackAESOneToMemory function
// This function is mainly used for unpack one file to memory.
// It will called by function UnpackAESToMemory.
func UnpackAESOneToMemory(data []byte, head TUnpackAESOne, dest *[]byte) (err error) {
	// initial, fill the key
	key := head.Key
	// first, split the data slice
	ss, err := SplitByte(data, AESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call AESDecrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go AESDecryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	r := bytes.Join(rr, []byte(""))
	// third, delete the more data
	size := BytesToInt(head.OriginSize)
	*dest = append(r[:0], r[:size]...)
	return err
}

// UnpackAESOneGo function
// This function is mainly used for unpack aes one file with go routine.
func UnpackAESOneGo(data []byte, head TUnpackAESOne, dest string, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	err = UnpackAESOne(data, head, dest)
	if err != nil {
		log.Println("Error aes unpack one:", err)
		return err
	}
	return err
}

// UnpackAESOneConfineGo function
// This function is mainly used for unpack aes one file with restrict go routine.
func UnpackAESOneConfineGo(data []byte, head TUnpackAESOne, dest string, wg *sync.WaitGroup, ch chan interface{}) (err error) {
	defer wg.Done()
	err = UnpackAESOneConfine(data, head, dest)
	if err != nil {
		log.Println("Error aes unpack one:", err)
		<-ch
		return err
	}
	<-ch
	return err
}

// UnpackAESOne function
// This function is mainly used for unpack aes one file.
func UnpackAESOne(data []byte, head TUnpackAESOne, path string) (err error) {
	// initial, fill the name
	var s []byte
	for _, v := range head.Name {
		if v == byte(0) {
			break
		}
		s = append(s, v)
	}
	file := path + string(s)
	key := head.Key
	// first, split the data slice
	ss, err := SplitByte(data, AESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call AESDecrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go AESDecryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// third, delete the more data
	size := BytesToInt(head.OriginSize)
	dest = append(dest[:0], dest[:size]...)
	// fourth, create the origin file
	err = ioutil.WriteFile(file, dest, 0644)
	if err != nil {
		log.Println("Error write to dest file:", err)
		return err
	}
	return err
}

// UnpackAESOneConfine function
// This function is mainly used for unpack aes one file with restrict go routine.
func UnpackAESOneConfine(data []byte, head TUnpackAESOne, path string) (err error) {
	// initial, fill the name
	var s []byte
	for _, v := range head.Name {
		if v == byte(0) {
			break
		}
		s = append(s, v)
	}
	file := path + string(s)
	key := head.Key
	// first, split the data slice
	ss, err := SplitByte(data, AESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call AESDecrypt function
	wg := &sync.WaitGroup{}
	ch := make(chan interface{}, ConfineBuffers)
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		ch <- struct{}{}
		go AESDecryptConfineGo(v, key, &rr[k], wg, ch)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// third, delete the more data
	size := BytesToInt(head.OriginSize)
	dest = append(dest[:0], dest[:size]...)
	// fourth, create the origin file
	err = ioutil.WriteFile(file, dest, 0644)
	if err != nil {
		log.Println("Error write to dest file:", err)
		return err
	}
	return err
}

func AESDecryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*dest, err = AESDecrypt(src, key)
	if err != nil {
		log.Println("Error aes decrypt data:", err)
		return err
	}
	atomic.AddInt64(&Done, 1)
	return err
}

func AESDecryptConfineGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup, ch chan interface{}) (err error) {
	defer wg.Done()
	*dest, err = AESDecrypt(src, key)
	if err != nil {
		log.Println("Error aes decrypt data:", err)
		<-ch
		return err
	}
	atomic.AddInt64(&Done, 1)
	<-ch
	return err
}

func AESDecrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 16, 24, 32
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error key length:", err)
		return dest, err
	}
	// calculate block size
	blockSize := block.BlockSize()
	// encrypt mode
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	// create slice
	dest = make([]byte, len(src))
	// decrypt
	blockMode.CryptBlocks(dest, src)
	// delete block data
	dest = PKCS7UnPadding(dest)
	return dest, err
}

func PKCS7UnPadding(src []byte) []byte {
	// don't need do anything
	/*size := len(src)
	unpadding := int(src[size-1])
	return src[:(size - unpadding)]*/
	return src
}
