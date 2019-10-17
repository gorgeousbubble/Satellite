package pack

import (
	"crypto/sha512"
	"encoding/hex"
	"sync"
)

func SHA512Check(src string, dest string) bool {
	b := false
	if SHA512Encode(src) == dest {
		b = true
	}
	return b
}

func SHA512Encode(src string) string {
	h := sha512.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA512EncryptGo(data []byte, r *[sha512.Size]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	*r = sha512.Sum512(data)
}

func SHA512Encrypt(data []byte) [sha512.Size]byte {
	return sha512.Sum512(data)
}
