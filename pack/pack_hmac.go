package pack

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// HMAC_SHA1 function
// input src string and key, output sha1 string
// src is a original string before sha1 calculate
// key is a string used by sha1
func HMAC_SHA1(src, key string) string {
	m := hmac.New(sha1.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// HMAC_SHA256 function
// input src string and key, output sha256 string
// src is a original string before sha256 calculate
// key is a string used by sha256
func HMAC_SHA256(src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// HMAC_SHA512 function
// input src string and key, output sha512 string
// src is a original string before sha512 calculate
// key is a string used by sha512
func HMAC_SHA512(src, key string) string {
	m := hmac.New(sha512.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}
