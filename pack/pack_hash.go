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
	case "BLAKE2B128", "blake2b128":
		dest = Blake2bEncode128(src)
	case "BLAKE2B256", "blake2b256":
		dest = Blake2bEncode256(src)
	case "BLAKE2B512", "blake2b512":
		dest = Blake2bEncode512(src)
	case "HMAC_SHA1", "hmac_sha1":
		dest = HMAC_SHA1(src, "")
	case "HMAC_SHA256", "hmac_sha256":
		dest = HMAC_SHA256(src, "")
	case "HMAC_SHA512", "hmac_sha512":
		dest = HMAC_SHA512(src, "")
	default:
		s := fmt.Sprint("Undefined hash algorithm.")
		err = errors.New(s)
	}
	return dest, err
}

func PackHashCheck(src string, dest string, algorithm string) (b bool, err error) {
	switch algorithm {
	case "MD5", "md5":
		b = MD5Check(src, dest)
	case "SHA1", "sha1":
		b = SHA1Check(src, dest)
	case "SHA256", "sha256":
		b = SHA256Check(src, dest)
	case "SHA512", "sha512":
		b = SHA512Check(src, dest)
	case "BLAKE2B128", "blake2b128":
		b = Blake2bCheck128(src, dest)
	case "BLAKE2B256", "blake2b256":
		b = Blake2bCheck256(src, dest)
	case "BLAKE2B512", "blake2b512":
		b = Blake2bCheck512(src, dest)
	default:
		s := fmt.Sprint("Undefined hash algorithm.")
		err = errors.New(s)
	}
	return b, err
}
