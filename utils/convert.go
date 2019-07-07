package utils

import (
	"bytes"
	"encoding/binary"
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