package pack

import (
	"crypto/sha1"
	"encoding/hex"
	"sync"
)

// SHA1Check function
// input src string and output dest string with sha1
// return bool indicate the success or failure function execute
func SHA1Check(src string, dest string) bool {
	b := false
	if SHA1Encode(src) == dest {
		b = true
	}
	return b
}

// SHA1Encode function
// encode src string output dest string with sha1
func SHA1Encode(src string) string {
	h := sha1.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1EncryptGo function
// encrypt sha1 with goroutine
func SHA1EncryptGo(data []byte, r *[sha1.Size]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	*r = sha1.Sum(data)
}

// SHA1Encrypt function
// encrypt sha1
func SHA1Encrypt(data []byte) [sha1.Size]byte {
	return sha1.Sum(data)
}
