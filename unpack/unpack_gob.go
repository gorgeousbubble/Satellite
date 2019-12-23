package unpack

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

func GobDecode(buf *bytes.Buffer, out interface{}) (err error) {
	dec := gob.NewDecoder(buf)
	err = dec.Decode(out)
	if err != nil {
		log.Println("Error gob decode interface:", err)
		return err
	}
	return err
}

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
