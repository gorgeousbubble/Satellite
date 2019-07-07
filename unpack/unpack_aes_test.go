package unpack

import (
	"io/ioutil"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	src := []byte{0x3B, 0x1B, 0x63, 0x41, 0x08, 0xC7, 0x8B, 0x97, 0xEC, 0x0D, 0xA3, 0xE4, 0xD2, 0xCD, 0x39, 0x84}
	key := []byte("Satellite-266414")
	r, err := AESDecrypt(src, key)
	if err != nil {
		t.Fatal("Error AES Decrypt:", err)
	}
	err = ioutil.WriteFile("test/file.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One:", err)
	}
}
