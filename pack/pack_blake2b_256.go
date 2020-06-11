package pack

import (
	"encoding/hex"
	"github.com/dchest/blake2b"
)

func Blake2bEncode256(src string) string {
	h := blake2b.New256()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}