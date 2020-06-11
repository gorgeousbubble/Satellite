package pack

import (
	"io/ioutil"
	"testing"
)

func TestBlake2bEncode256(t *testing.T) {
	src := "hello,world!"
	r := Blake2bEncode256(src)
	err := ioutil.WriteFile("../test/data/pack/file_blake2b_256.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write SHA1 Encode:", err)
	}
}
