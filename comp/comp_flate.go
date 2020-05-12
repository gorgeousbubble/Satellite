package comp

import (
	"bytes"
	"compress/flate"
	"log"
)

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
