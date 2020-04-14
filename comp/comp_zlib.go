package comp

import (
	"bytes"
	"compress/zlib"
	"log"
)

// CompressZlib function
// input src file list, output dest file path, return error info
// this function will use zlib algorithm to compress file list
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// return err indicate the success or failure function execute
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
