package pack

import (
	"io/ioutil"
	. "satellite/utils"
	"testing"
)

func TestRSAEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	r, err := RSAEncrypt(src, ConstRSAPublicKey)
	if err != nil {
		t.Fatal("Error RSA Encrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write RSA Encrypt:", err)
	}
}

func BenchmarkRSAEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r, err := RSAEncrypt(src, ConstRSAPublicKey)
		if err != nil {
			b.Fatal("Error RSA Encrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write RSA Encrypt:", err)
		}
	}
}
