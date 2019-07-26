package unpack

import (
	"io/ioutil"
	"testing"
)

func TestDESDecrypt(t *testing.T) {
	src := []byte{0x8F, 0x34, 0x1C, 0x84, 0xBB, 0xB4, 0x58, 0xA8, 0x54, 0x8B, 0xC0, 0x9E, 0x2F, 0x55, 0x01, 0xF2}
	key := []byte("hyacinth")
	r, err := DESDecrypt(src, key)
	if err != nil {
		t.Fatal("Error DES Decrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES One:", err)
	}
}

func BenchmarkDESDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{0x8F, 0x34, 0x1C, 0x84, 0xBB, 0xB4, 0x58, 0xA8, 0x54, 0x8B, 0xC0, 0x9E, 0x2F, 0x55, 0x01, 0xF2}
		key := []byte("hyacinth")
		r, err := DESDecrypt(src, key)
		if err != nil {
			b.Fatal("Error DES Decrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES One:", err)
		}
	}
}
