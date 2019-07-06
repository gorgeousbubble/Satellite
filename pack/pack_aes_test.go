package pack

import (
	"io/ioutil"
	"testing"
)

func TestPackAESOne(t *testing.T) {
	src := "test/file.txt"
	r, err := PackAESOne(src)
	if err != nil {
		t.Fatal("Error Pack AES One:", err)
	}
	err = ioutil.WriteFile("test/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One:", err)
	}
}

func TestAESEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	key := []byte("Satellite-266414")
	_, err := AESEncrypt(src, key)
	if err != nil {
		t.Fatal("Error AES Encrypt:", err)
	}
}

func BenchmarkPackAESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "test/file.txt"
		r, err := PackAESOne(src)
		if err != nil {
			b.Fatal("Error Pack AES One:", err)
		}
		err = ioutil.WriteFile("test/file_aes.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write AES One:", err)
		}
	}
}

func BenchmarkAESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		key := []byte("Satellite-266414")
		_, err := AESEncrypt(src, key)
		if err != nil {
			b.Fatal("Error AES Encrypt:", err)
		}
	}
}