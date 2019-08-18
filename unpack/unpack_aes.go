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
)

func UnpackAES(srcfile string, destpath string) (err error) {
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
		go UnpackAESOneGo(s, hh, destpath, wg)
	}
	wg.Wait()
	return err
}

func UnpackAESOneGo(data []byte, head TUnpackAESOne, destpath string, wg *sync.WaitGroup) (err error) {
	err = UnpackAESOne(data, head, destpath)
	if err != nil {
		log.Println("Error AES Unpack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func UnpackAESOne(data []byte, head TUnpackAESOne, destpath string) (err error) {
	// initial, fill the name
	var s []byte
	for _, v := range head.Name {
		if v == byte(0) {
			break
		}
		s = append(s, v)
	}
	destfile := destpath + string(s)
	key := head.Key
	// first, split the data slice
	ss, err := SplitByte(data, ConstAESBufferSize)
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
	err = ioutil.WriteFile(destfile, dest, 0644)
	if err != nil {
		log.Println("Error write to destfile:", err)
		return err
	}
	return err
}

func AESDecryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = AESDecrypt(src, key)
	if err != nil {
		log.Println("Error AES Decrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
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
