package pack

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

// GobEncode function
// Use gob encode buffer stream
// input interface and output buffer stream
// return err indicate the success or failure function execute
func GobEncode(buf *bytes.Buffer, in interface{}) (err error) {
	enc := gob.NewEncoder(buf)
	err = enc.Encode(in)
	if err != nil {
		log.Println("Error gob encode interface:", err)
		return err
	}
	return err
}

// GobEncodeTo function
// Use gob encode to file
// input interface and output to file
// return err indicate the success or failure function execute
func GobEncodeTo(src string, in interface{}) (err error) {
	file, err := os.OpenFile(src, os.O_CREATE|os.O_WRONLY, 0)
	if err != nil {
		log.Println("Error open gob file:", err)
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	err = enc.Encode(in)
	if err != nil {
		log.Println("Error gob encode interface:", err)
		return err
	}
	return err
}
