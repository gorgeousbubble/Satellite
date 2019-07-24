package pack

import (
	"crypto/sha512"
	"sync"
)

func SHA512EncryptGo(data []byte, r *[sha512.Size]byte, wg *sync.WaitGroup) {
	*r = sha512.Sum512(data)
	wg.Done()
}

func SHA512Encrypt(data []byte) [sha512.Size]byte {
	return sha512.Sum512(data)
}
