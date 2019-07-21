package unpack

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"sync"
)

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
