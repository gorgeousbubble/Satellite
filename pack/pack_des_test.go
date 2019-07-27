package pack

import (
	"io/ioutil"
	"sync"
	"testing"
)

func TestPackDES(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_des.txt"
	err := PackDES(src, dest)
	if err != nil {
		t.Fatal("Error Pack DES:", err)
	}
}

func TestPack3DESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := "../test/data/pack/file.txt"
	wg.Add(1)
	err := Pack3DESOneGo(src, &r, &wg)
	if err != nil {
		t.Fatal("Error Pack 3DES One Go:", err)
	}
	wg.Wait()
	err = ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write 3DES One Go:", err)
	}
}

func TestPack3DESOne(t *testing.T) {
	src := "../test/data/pack/file.txt"
	r, err := Pack3DESOne(src)
	if err != nil {
		t.Fatal("Error Pack 3DES One:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write 3DES One:", err)
	}
}

func TestPackDESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := "../test/data/pack/file.txt"
	wg.Add(1)
	err := PackDESOneGo(src, &r, &wg)
	if err != nil {
		t.Fatal("Error Pack DES One Go:", err)
	}
	wg.Wait()
	err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES One Go:", err)
	}
}

func TestPackDESOne(t *testing.T) {
	src := "../test/data/pack/file.txt"
	r, err := PackDESOne(src)
	if err != nil {
		t.Fatal("Error Pack DES One:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES One:", err)
	}
}

func TestTripleDESEncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := []byte("hello,world!")
	key := []byte("HyacinthRaindropRomantic")
	wg.Add(1)
	go TripleDESEncryptGo(src, key, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write 3DES Encrypt:", err)
	}
}

func TestTripleDESEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	key := []byte("HyacinthRaindropRomantic")
	r, err := TripleDESEncrypt(src, key)
	if err != nil {
		t.Fatal("Error 3DES Encrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write 3DES Encrypt:", err)
	}
}

func TestDESEncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := []byte("hello,world!")
	key := []byte("hyacinth")
	wg.Add(1)
	go DESEncryptGo(src, key, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES Encrypt:", err)
	}
}

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

func BenchmarkPackDES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_des.txt"
		err := PackDES(src, dest)
		if err != nil {
			b.Fatal("Error Pack DES:", err)
		}
	}
}

func BenchmarkPack3DESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := "../test/data/pack/file.txt"
		wg.Add(1)
		err := Pack3DESOneGo(src, &r, &wg)
		if err != nil {
			b.Fatal("Error Pack 3DES One Go:", err)
		}
		wg.Wait()
		err = ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write 3DES One Go:", err)
		}
	}
}

func BenchmarkPack3DESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/pack/file.txt"
		r, err := Pack3DESOne(src)
		if err != nil {
			b.Fatal("Error Pack 3DES One:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write 3DES One:", err)
		}
	}
}

func BenchmarkPackDESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := "../test/data/pack/file.txt"
		wg.Add(1)
		err := PackDESOneGo(src, &r, &wg)
		if err != nil {
			b.Fatal("Error Pack DES One Go:", err)
		}
		wg.Wait()
		err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES One Go:", err)
		}
	}
}

func BenchmarkPackDESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/pack/file.txt"
		r, err := PackDESOne(src)
		if err != nil {
			b.Fatal("Error Pack DES One:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES One:", err)
		}
	}
}

func BenchmarkTripleDESEncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := []byte("hello,world!")
		key := []byte("HyacinthRaindropRomantic")
		wg.Add(1)
		go TripleDESEncryptGo(src, key, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write 3DES Encrypt:", err)
		}
	}
}

func BenchmarkTripleDESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		key := []byte("HyacinthRaindropRomantic")
		r, err := TripleDESEncrypt(src, key)
		if err != nil {
			b.Fatal("Error 3DES Encrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_3des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write 3DES Encrypt:", err)
		}
	}
}

func BenchmarkDESEncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := []byte("hello,world!")
		key := []byte("hyacinth")
		wg.Add(1)
		go DESEncryptGo(src, key, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES Encrypt:", err)
		}
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
