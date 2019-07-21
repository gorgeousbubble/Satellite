package pack

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"sync"
)

func RSAEncryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = RSAEncrypt(src, key)
	if err != nil {
		log.Println("Error RSA Encrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func RSAEncrypt(src, key []byte) (dest []byte, err error) {
	block, _ := pem.Decode(key)
	if block == nil {
		err = errors.New("RSA Public Key Error")
		return dest, err
	}
	pi, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return dest, err
	}
	pub := pi.(*rsa.PublicKey)
	dest, err = rsa.EncryptPKCS1v15(rand.Reader, pub, src)
	if err != nil {
		return dest, err
	}
	return dest, err
}