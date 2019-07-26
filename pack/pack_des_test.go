package pack

import (
	"io/ioutil"
	"testing"
)

func TestDESEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	key := []byte("hyacinth")
	r, err := DESEncrypt(src, key)
	if err != nil {
		t.Fatal("Error DES Encrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES Encrypt:", err)
	}
}

func BenchmarkDESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		key := []byte("hyacinth")
		r, err := DESEncrypt(src, key)
		if err != nil {
			b.Fatal("Error DES Encrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES Encrypt:", err)
		}
	}
}
