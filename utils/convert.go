package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToBytes(n int) []byte {
	x := int32(n)
	r := bytes.NewBuffer([]byte{})
	binary.Write(r, binary.BigEndian, x)
	return r.Bytes()
}

func BytesToInt(b []byte) int {
	var x int32
	r := bytes.NewBuffer(b)
	binary.Read(r, binary.BigEndian, &x)
	return int(x)
}

func BytesCopy(r *[]byte, s []byte) bool {
	if len(*r) < len(s) {
		return false
	}
	for i := 0; i < len(s); i++ {
		(*r)[i] = s[i]
	}
	return true
}

func SplitByte(data []byte, size int) (r [][]byte, err error) {
	rd := bytes.NewReader(data)
	for {
		s := make([]byte, size)
		switch n, err := rd.Read(s); true {
		case n < 0:
			log.Println("Error read byte:", err)
			return r, err
		case n == 0:
			return r, nil
		case n > 0:
			r = append(r, s)
		}
	}
}