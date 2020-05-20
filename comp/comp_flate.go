package comp

import (
	"bytes"
	"compress/flate"
	"log"
)

// CompressGzip function
// input src byte stream, output dest byte stream, return error info
// this function will use flate algorithm to compress stream
// src stream should be byte slice stream, if input a string, please convert it first
// dest stream should be byte slice stream, if input a string, please convert it first
// return err indicate the success or failure function execute
func CompressFlate(src []byte) (dest []byte, err error) {
	// create new buffer
	out := bytes.NewBuffer(nil)
	// create new flate writer
	fw, err := flate.NewWriter(out, flate.BestCompression)
	if err != nil {
		log.Println("Error create new flate writer:", err)
		return dest, err
	}
	defer fw.Close()
	// write stream to buffer
	_, err = fw.Write(src)
	if err != nil {
		log.Println("Error write stream to buffer:", err)
		return dest, err
	}
	// flush the buffer
	err = fw.Flush()
	if err != nil {
		log.Println("Error flush buffer:", err)
		return dest, err
	}
	dest = out.Bytes()
	return dest, err
}
