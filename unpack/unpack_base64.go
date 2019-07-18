package unpack

import "encoding/base64"

func Base64Decrypt(str string) string {
	s, _ := base64.StdEncoding.DecodeString(str)
	r := string(s)
	return r
}
