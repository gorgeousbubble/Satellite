package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

var RSAPrivateKey []byte
var RSAPublicKey []byte

func init() {
	err := GenRSAKey2Memory(&RSAPrivateKey, &RSAPublicKey, 1024)
	if err != nil {
		log.Println("Error Gen RSA Key:", err)
	}
}

func GenRSAKey2File(bits int) error {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derSteam := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: derSteam,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// generate public key
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil
	}
	block = &pem.Block{
		Type: "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func GenRSAKey2Memory(pri *[]byte, pub *[]byte, bits int) error {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derSteam := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: derSteam,
	}
	*pri = pem.EncodeToMemory(block)
	// generate public key
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil
	}
	block = &pem.Block{
		Type: "PUBLIC KEY",
		Bytes: derPkix,
	}
	*pub = pem.EncodeToMemory(block)
	return nil
}
