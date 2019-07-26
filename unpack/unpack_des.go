package unpack

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"io/ioutil"
	"log"
	. "satellite/utils"
	"sync"
)

func UnpackDESOneGo(data []byte, head TUnpackDESOne, destpath string, wg *sync.WaitGroup) (err error) {
	err = UnpackDESOne(data, head, destpath)
	if err != nil {
		log.Println("Error DES Unpack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func UnpackDESOne(data []byte, head TUnpackDESOne, destpath string) (err error) {
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
	ss, err := SplitByte(data, ConstDESBufferSize)
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
	err = ioutil.WriteFile(destfile, dest, 0644)
	if err != nil {
		log.Println("Error write to destfile:", err)
		return err
	}
	return err
}

func DESDecryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = DESDecrypt(src, key)
	if err != nil {
		log.Println("Error DES Decrypt data:", err)
		wg.Done()
		return err
	}
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
	blockMode := cipher.NewCBCDecrypter(block,key[:blockSize])
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
