package unpack

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
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

func Unpack3DES(src string, dest string) (err error) {
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
	h := TUnpackDES{}
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
	BytesCopy(&s, []byte("3DES"))
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
		hh := TUnpack3DESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 24)
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
		go Unpack3DESOneGo(s, hh, dest, wg)
	}
	wg.Wait()
	return err
}

func UnpackDES(src string, dest string) (err error) {
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
	h := TUnpackDES{}
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
	BytesCopy(&s, []byte("DES"))
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
		hh := TUnpackDESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 8)
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
		go UnpackDESOneGo(s, hh, dest, wg)
	}
	wg.Wait()
	return err
}

func Unpack3DESExtractInfo(src string, dest *[]string, sz *[]int) (err error) {
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
	h := TUnpackDES{}
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
	BytesCopy(&s, []byte("3DES"))
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
		hh := TUnpack3DESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 24)
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
		*dest = append(*dest, string(hh.Name))
		*sz = append(*sz, BytesToInt(hh.OriginSize))
	}
	return err
}

func UnpackDESExtractInfo(src string, dest *[]string, sz *[]int) (err error) {
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
	h := TUnpackDES{}
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
	BytesCopy(&s, []byte("DES"))
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
		hh := TUnpackDESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 8)
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
		*dest = append(*dest, string(hh.Name))
		*sz = append(*sz, BytesToInt(hh.OriginSize))
	}
	return err
}

func Unpack3DESWorkCalculate(src string) (work int64, err error) {
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
	h := TUnpackDES{}
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
	BytesCopy(&s, []byte("3DES"))
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
		hh := TUnpack3DESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 24)
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

func UnpackDESWorkCalculate(src string) (work int64, err error) {
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
	h := TUnpackDES{}
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
	BytesCopy(&s, []byte("DES"))
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
		hh := TUnpackDESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 8)
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

func Unpack3DESOneToMemory(data []byte, head TUnpack3DESOne, dest *[]byte) (err error) {
	// initial, fill the key
	key := head.Key
	// first, split the data slice
	ss, err := SplitByte(data, DESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call DESDecrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go TripleDESDecryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	r := bytes.Join(rr, []byte(""))
	// third, delete the more data
	size := BytesToInt(head.OriginSize)
	*dest = append(r[:0], r[:size]...)
	return err
}

func Unpack3DESOneGo(data []byte, head TUnpack3DESOne, dest string, wg *sync.WaitGroup) (err error) {
	err = Unpack3DESOne(data, head, dest)
	if err != nil {
		log.Println("Error 3des unpack one file:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func Unpack3DESOne(data []byte, head TUnpack3DESOne, path string) (err error) {
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
	ss, err := SplitByte(data, DESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call DESDecrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go TripleDESDecryptGo(v, key, &rr[k], wg)
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

func UnpackDESOneToMemory(data []byte, head TUnpackDESOne, dest *[]byte) (err error) {
	// initial, fill the key
	key := head.Key
	// first, split the data slice
	ss, err := SplitByte(data, DESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call DESDecrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go DESDecryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	r := bytes.Join(rr, []byte(""))
	// third, delete the more data
	size := BytesToInt(head.OriginSize)
	*dest = append(r[:0], r[:size]...)
	return err
}

func UnpackDESOneGo(data []byte, head TUnpackDESOne, dest string, wg *sync.WaitGroup) (err error) {
	err = UnpackDESOne(data, head, dest)
	if err != nil {
		log.Println("Error des unpack one:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func UnpackDESOne(data []byte, head TUnpackDESOne, path string) (err error) {
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
	ss, err := SplitByte(data, DESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call DESDecrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go DESDecryptGo(v, key, &rr[k], wg)
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

func TripleDESDecryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = TripleDESDecrypt(src, key)
	if err != nil {
		log.Println("Error 3des decrypt data:", err)
		wg.Done()
		return err
	}
	atomic.AddInt64(&Done, 1)
	wg.Done()
	return err
}

func TripleDESDecrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 24
	block, err := des.NewTripleDESCipher(key)
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
	dest = PKCS5UnPadding(dest)
	return dest, err
}

func DESDecryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = DESDecrypt(src, key)
	if err != nil {
		log.Println("Error des decrypt data:", err)
		wg.Done()
		return err
	}
	atomic.AddInt64(&Done, 1)
	wg.Done()
	return err
}

func DESDecrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 8
	block, err := des.NewCipher(key)
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
	dest = PKCS5UnPadding(dest)
	return dest, err
}

func PKCS5UnPadding(src []byte) []byte {
	// don't need do anything
	/*size := len(src)
	unpadding := int(src[size-1])
	return src[:(size - unpadding)]*/
	return src
}
