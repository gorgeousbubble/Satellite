package pack

import (
	"crypto/sha512"
	"encoding/hex"
	"sync"
)

// SHA512Check function
// input src string and output dest string with sha512
// return bool indicate the success or failure function execute
func SHA512Check(src string, dest string) bool {
	b := false
	if SHA512Encode(src) == dest {
		b = true
	}
	return b
}

// SHA512Encode function
// encode src string output dest string with sha512
func SHA512Encode(src string) string {
	h := sha512.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA512EncryptGo function
// encrypt sha512 with goroutine
func SHA512EncryptGo(data []byte, r *[sha512.Size]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	*r = sha512.Sum512(data)
}

// SHA512Encrypt function
// encrypt sha512
func SHA512Encrypt(data []byte) [sha512.Size]byte {
	return sha512.Sum512(data)
}
