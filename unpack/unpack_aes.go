package unpack

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"sync"
)

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
	dest = []byte{}
	err = nil
	// key length should be 16, 24, 32
	block, err := aes.NewCipher(key)
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
	dest = PKCS7UnPadding(dest)
	return dest, err
}

func PKCS7UnPadding(src []byte) []byte {
	size := len(src)
	unpadding := int(src[size-1])
	return src[:(size - unpadding)]
}