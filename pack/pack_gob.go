package pack

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

func GobEncodeStream(buf *bytes.Buffer, in interface{}) (err error) {
	enc := gob.NewEncoder(buf)
	err = enc.Encode(in)
	if err != nil {
		log.Println("Error gob encode interface:", err)
		return err
	}
	return err
}

func GobEncodeFile(src string, in interface{}) (err error) {
	file, err := os.Open(src)
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
