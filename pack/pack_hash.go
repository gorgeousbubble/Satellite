package pack

import (
	"errors"
	"fmt"
)

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
