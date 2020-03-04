package pack

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

// MD5Check function
// check the string before and after MD5 calculate equal
// input src string and dest string, compare whether they are equal
// return bool indicate true or false two string equal with MD5
func MD5Check(src string, dest string) bool {
	b := false
	if MD5Encode(src) == dest {
		b = true
	}
	return b
}

// MD5Encode function
// encode MD5
// input src string and output string after MD5 encode
// return string after MD5 encode
func MD5Encode(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5EncryptGo function
func MD5EncryptGo(data []byte, r *[md5.Size]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	*r = md5.Sum(data)
}

// MD5Encrypt function
func MD5Encrypt(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}
