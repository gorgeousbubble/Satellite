package unpack

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	. "satellite/utils"
)

func RSADecrypt(src []byte) (dest []byte, err error) {
	block, _ := pem.Decode(RSAPrivateKey)
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
