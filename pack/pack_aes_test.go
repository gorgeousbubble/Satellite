package pack

import (
	"io/ioutil"
	. "satellite/utils"
	"sync"
	"testing"
)

func TestPackAES(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	err := PackAES(src, dest)
	if err != nil {
		t.Fatal("Error Pack AES:", err)
	}
}

func TestPackAESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := "../test/data/pack/file.txt"
	wg.Add(1)
	err := PackAESOneGo(src, &r, &wg)
	if err != nil {
		t.Fatal("Error Pack AES One Go:", err)
	}
	wg.Wait()
	err = ioutil.WriteFile("../test/data/pack/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One Go:", err)
	}
}

func TestPackAESOne(t *testing.T) {
	src := "../test/data/pack/file.txt"
	r, err := PackAESOne(src)
	if err != nil {
		t.Fatal("Error Pack AES One:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One:", err)
	}
}

func TestAESEncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := []byte("hello,world!")
	key := []byte("Satellite-266414")
	wg.Add(1)
	go AESEncryptGo(src, key, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES Encrypt:", err)
	}
}

func TestAESEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	key := []byte("Satellite-266414")
	r, err := AESEncrypt(src, key)
	if err != nil {
		t.Fatal("Error AES Encrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES Encrypt:", err)
	}
}

func TestSplitByte(t *testing.T) {
	data := []byte("hello,World!")
	size := ConstAESBufferSize
	_, err := SplitByte(data, size)
	if err != nil {
		t.Fatal("Error Split Byte:", err)
	}
}

func BenchmarkPackAES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_aes.txt"
		err := PackAES(src, dest)
		if err != nil {
			b.Fatal("Error Pack AES:", err)
		}
	}
}

func BenchmarkPackAESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := "../test/data/pack/file.txt"
		wg.Add(1)
		err := PackAESOneGo(src, &r, &wg)
		if err != nil {
			b.Fatal("Error Pack AES One Go:", err)
		}
		wg.Wait()
		err = ioutil.WriteFile("../test/data/pack/file_aes.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write AES One Go:", err)
		}
	}
}

func BenchmarkPackAESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/pack/file.txt"
		r, err := PackAESOne(src)
		if err != nil {
			b.Fatal("Error Pack AES One:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_aes.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write AES One:", err)
		}
	}
}

func BenchmarkAESEncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := []byte("hello,world!")
		key := []byte("Satellite-266414")
		wg.Add(1)
		go AESEncryptGo(src, key, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_aes.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write AES Encrypt:", err)
		}
	}
}

func BenchmarkAESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		key := []byte("Satellite-266414")
		_, err := AESEncrypt(src, key)
		if err != nil {
			b.Fatal("Error Write AES Encrypt:", err)
		}
	}
}

func BenchmarkSplitByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []byte("hello,World!")
		size := ConstAESBufferSize
		_, err := SplitByte(data, size)
		if err != nil {
			b.Fatal("Error Split Byte:", err)
		}
	}
}