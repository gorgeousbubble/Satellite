package unpack

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

// GobDecode function
// This function is mainly used for decode gob type
// buf buffer is data stream wait for decode
// out interface should input structure address
// return err indicate the success or failure function execute
func GobDecode(buf *bytes.Buffer, out interface{}) (err error) {
	dec := gob.NewDecoder(buf)
	err = dec.Decode(out)
	if err != nil {
		log.Println("Error gob decode interface:", err)
		return err
	}
	return err
}

// GobDecodeFrom function
// This function is mainly used for decode gob type from file
// src file save gob data wait for decode
// out interface should input structure address
// return err indicate the success or failure function execute
func GobDecodeFrom(src string, out interface{}) (err error) {
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open gob file:", err)
		return err
	}
	defer file.Close()
	dec := gob.NewDecoder(file)
	err = dec.Decode(out)
	if err != nil {
		log.Println("Error gob decode interface:", err)
		return err
	}
	return err
}
