package unpack

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	. "satellite/utils"
	"sync"
)

func UnpackRSAOneGo(data []byte, head TUnpackRSAOne, destpath string, wg *sync.WaitGroup) (err error) {
	err = UnpackRSAOne(data, head, destpath)
	if err != nil {
		log.Println("Error RSA Unpack One:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func UnpackRSAOne(data []byte, head TUnpackRSAOne, destpath string) (err error) {
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
	ss, err := SplitByte(data, ConstRSAUnpackSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	// second, we can call RSADecrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go RSADecryptGo(v, key, &rr[k], wg)
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

func RSADecryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = RSADecrypt(src, key)
	if err != nil {
		log.Println("Error RSA Decrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func RSADecrypt(src, key []byte) (dest []byte, err error) {
	block, _ := pem.Decode(key)
	if block == nil {
		err = errors.New("RSA Private Key Error")
		return dest, err
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return dest, err
	}
	dest, err = rsa.DecryptPKCS1v15(rand.Reader, pri, src)
	if err != nil {
		return dest, err
	}
	return dest, err
}
