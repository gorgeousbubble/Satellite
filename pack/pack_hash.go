package pack

import (
	"errors"
	"fmt"
)

// PackHashEncode function
// input src string, output hash string and algorithm which used in hash, return error info
// this function will base on algorithm to call correspond function
// src string which you want to encode by hash algorithm, like 'hello,world!' or '../test/data/file.txt'
// dest string is the result of hash value, like 'C:\\package.pak' or '../test/data/package.pak'
// algorithm now support 'md5', 'sha1', 'sha256', 'sha512', you can send both up case and low case
// return err indicate the success or failure function execute
func PackHashEncode(src string, algorithm string) (dest string, err error) {
	switch algorithm {
	case "MD5", "md5":
		dest = MD5Encode(src)
	case "SHA1", "sha1":
		dest = SHA1Encode(src)
	case "SHA256", "sha256":
		dest = SHA256Encode(src)
	case "SHA512", "sha512":
		dest = SHA512Encode(src)
	default:
		s := fmt.Sprint("Undefined hash algorithm.")
		err = errors.New(s)
	}
	return dest, err
}
