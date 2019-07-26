package pack

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"log"
)

func DESEncrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 8
	block, err := des.NewCipher(key)
	if err != nil {
		log.Println("Error key length:", err)
		return dest, err
	}
	// calculate block size
	blockSize := block.BlockSize()
	// fill block data
	src = PKCS5Padding(src, blockSize)
	// encrypt mode
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// create slice
	dest = make([]byte, len(src))
	// encrypt
	blockMode.CryptBlocks(dest, src)
	return dest, err
}

func PKCS5Padding(src []byte, size int) []byte {
	var padding int
	if len(src)%size != 0 {
		padding = size - len(src) % size
	}
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, text...)
}
