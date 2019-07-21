package pack

import (
	"crypto/md5"
	"sync"
)

func MD5EncryptGo(data []byte, r *[16]byte, wg *sync.WaitGroup) {
	*r = md5.Sum(data)
	wg.Done()
}

func MD5Encrypt(data []byte) [16]byte {
	return md5.Sum(data)
}
