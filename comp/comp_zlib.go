package comp

import (
	"bytes"
	"compress/zlib"
	"log"
)

func CompressZlib(src []byte) (dest []byte, err error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err = w.Write(src)
	if err != nil {
		log.Println("Error compress zlib:", err)
		return dest, err
	}
	err = w.Close()
	if err != nil {
		log.Println("Error close zlib writer:", err)
		return dest, err
	}
	dest = in.Bytes()
	return dest, err
}
