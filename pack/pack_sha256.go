package pack

import "crypto/sha256"

func SHA256Encrypt(data []byte) [sha256.Size]byte {
	return sha256.Sum256(data)
}
