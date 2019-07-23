package pack

import (
	"crypto/sha256"
	"sync"
)

func SHA256EncryptGo(data []byte, r *[sha256.Size]byte, wg *sync.WaitGroup) {
	*r = sha256.Sum256(data)
	wg.Done()
}

func SHA256Encrypt(data []byte) [sha256.Size]byte {
	return sha256.Sum256(data)
}
