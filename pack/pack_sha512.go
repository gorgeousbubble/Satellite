package pack

import "crypto/sha512"

func SHA512Encrypt(data []byte) [sha512.Size]byte {
	return sha512.Sum512(data)
}
