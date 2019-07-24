package pack

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

func MD5Check(src string, dest string) bool {
	b := false
	if MD5Encode(src) == dest {
		b = true
	}
	return b
}

func MD5Encode(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5EncryptGo(data []byte, r *[md5.Size]byte, wg *sync.WaitGroup) {
	*r = md5.Sum(data)
	wg.Done()
}

func MD5Encrypt(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}
