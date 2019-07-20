package unpack

import (
	"encoding/base64"
	"sync"
)

func Base64DecryptGo(str string, r *string, wg *sync.WaitGroup) {
	*r = Base64Decrypt(str)
	wg.Done()
}

func Base64Decrypt(str string) string {
	s, _ := base64.StdEncoding.DecodeString(str)
	r := string(s)
	return r
}
