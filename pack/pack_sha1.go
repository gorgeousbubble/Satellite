package pack

import (
	"crypto/sha1"
	"sync"
)

func SHA1EncryptGo(data []byte, r *[sha1.Size]byte, wg *sync.WaitGroup)  {
	*r = sha1.Sum(data)
	wg.Done()
}

func SHA1Encrypt(data []byte) [sha1.Size]byte {
	return sha1.Sum(data)
}
