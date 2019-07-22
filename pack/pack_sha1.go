package pack

import "crypto/sha1"

func SHA1Encrypt(data []byte) [sha1.Size]byte {
	return sha1.Sum(data)
}
