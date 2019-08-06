package decomp

import (
	"bytes"
	"compress/zlib"
	"io"
	"log"
)

func DeCompressZlib(src []byte) (dest []byte, err error) {
	var out bytes.Buffer
	b := bytes.NewReader(src)
	r, err := zlib.NewReader(b)
	if err != nil {
		log.Println("Error decompress zlib:", err)
		return dest, err
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		log.Println("Error writer zlib:", err)
		return dest, err
	}
	dest = out.Bytes()
	return dest, err
}
