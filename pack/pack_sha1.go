package pack

import (
	"crypto/sha1"
	"encoding/hex"
	"sync"
)

func SHA1Check(src string, dest string) bool {
	b := false
	if SHA1Encode(src) == dest {
		b = true
	}
	return b
}

func SHA1Encode(src string) string {
	h := sha1.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1EncryptGo(data []byte, r *[sha1.Size]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	*r = sha1.Sum(data)
}

func SHA1Encrypt(data []byte) [sha1.Size]byte {
	return sha1.Sum(data)
}
