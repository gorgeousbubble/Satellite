package pack

import (
	"bytes"
	"encoding/gob"
	"log"
)

func GobEncodeStream(buf *bytes.Buffer, e interface{}) (err error) {
	enc := gob.NewEncoder(buf)
	err = enc.Encode(e)
	if err != nil {
		log.Println("Error gob encode interface:", err)
		return err
	}
	return err
}
