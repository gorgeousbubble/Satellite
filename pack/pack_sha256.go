package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

func SHA256Check(src string, dest string) bool {
	b := false
	if SHA256Encode(src) == dest {
		b = true
	}
	return b
}

func SHA256Encode(src string) string {
	h := sha256.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256EncryptGo(data []byte, r *[sha256.Size]byte, wg *sync.WaitGroup) {
	*r = sha256.Sum256(data)
	wg.Done()
}

func SHA256Encrypt(data []byte) [sha256.Size]byte {
	return sha256.Sum256(data)
}
