package pack

import (
	"crypto/md5"
	"sync"
)

func MD5EncryptGo(data []byte, r *[md5.Size]byte, wg *sync.WaitGroup) {
	*r = md5.Sum(data)
	wg.Done()
}

func MD5Encrypt(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}
