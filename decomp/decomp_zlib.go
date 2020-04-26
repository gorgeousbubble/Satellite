package decomp

import (
	"bytes"
	"compress/zlib"
	"io"
	"log"
)

// DeCompressZlib function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src data wait for decompress by zlib
// dest date result decompress by zlib
// return err indicate the success or failure function execute
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
