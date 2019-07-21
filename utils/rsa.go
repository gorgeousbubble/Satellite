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

var ConstRSAPrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQD0z8923lo7lQTIBZmmtPPWjCj+TyEMLoFh3NBZ9r0u6zVs5iPy
TVPguWQi1q5UOD1uZCk8cJN+zlfg0VOE3DBmYGsDDYm3E5KbIUe/xSTPFSO1K3Hc
jPHH+zsWBfGuR1kBbwz3Y5udTyxT21ToItIBBZsv84+kPdrW4Ig+GKbK7wIDAQAB
AoGBAJ4MQ34UUuDAdhPEOcw8ameKmTSFVWqN442QhxpthvlxdE5XzcEyVwJv7cvo
GSfaHx7TnyOb8j0dbfKAcZEGrl9GfdH/q1qtNNCo/xorMwkvbe9XXaaZp2g/pbQr
ew1FZuNXrNo4m2n0P3MpDq85BJeImKcYmXbVzwedWyVOML5BAkEA/dpEhROPCJ2m
QCV/xCNMLgPlV9QmpKa+Z230FTaqUZ5f1xk1Jmv/fOYLV15Jp8kma6IZEc3ZhcgJ
LtJaxuba0QJBAPbh9trcvGAaiycap9FkhUvJej1HAYggCubvwU4xRomnEj1l5cJZ
b0hJ5Jrxw8pPdfkfj5I8pXCUo+ssntrhOb8CQEmNbZ8VbLM1Yo9hjiSZiaAnltMc
8FlyyuaTEE75ON5PQjvD3QeV8UASM6UL99F60fwLyrHC8Ez4CdkcMBtDlpECQQDF
rVqWFeZzYO2IqwwyWjVs4G8TP5aXY2i3TsPOEyByuaaeKMfWXPwAusHj2q81f3hr
t/yiUBL72NiDRqAPjAVLAkABgVa67I5xPpC6JAXgTwXrVylSW7ORv9Gpy2icL6Zo
M+f6JxdroXs4pTGru68Buv+Okemt09djHyIahm+BNOkA
-----END RSA PRIVATE KEY-----
`)

var ConstRSAPublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD0z8923lo7lQTIBZmmtPPWjCj+
TyEMLoFh3NBZ9r0u6zVs5iPyTVPguWQi1q5UOD1uZCk8cJN+zlfg0VOE3DBmYGsD
DYm3E5KbIUe/xSTPFSO1K3HcjPHH+zsWBfGuR1kBbwz3Y5udTyxT21ToItIBBZsv
84+kPdrW4Ig+GKbK7wIDAQAB
-----END PUBLIC KEY-----
`)

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
