package pack

import (
	"encoding/hex"
	"github.com/dchest/blake2b"
)

func Blake2bEncode128(src string) string {
	h := blake2b.NewMAC(16, nil)
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func Blake2bEncode256(src string) string {
	h := blake2b.New256()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func Blake2bEncode512(src string) string {
	h := blake2b.New512()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}
