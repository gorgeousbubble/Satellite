package pack

import (
	"io/ioutil"
	"sync"
	"testing"
)

// TestPack3DES function
func TestPack3DES(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_3des.txt"
	err := Pack3DES(src, dest)
	if err != nil {
		t.Fatal("Error Pack DES:", err)
	}
}

// TestPackDES function
func TestPackDES(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_des.txt"
	err := PackDES(src, dest)
	if err != nil {
		t.Fatal("Error Pack DES:", err)
	}
}

// TestPackDESWorkCalculate function
func TestPackDESWorkCalculate(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	_, err := PackDESWorkCalculate(src)
	if err != nil {
		t.Fatal("Error Pack DES Work Calculate:", err)
	}
}

// TestPack3DESOneGo function
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

// TestPack3DESOne function
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

// TestPackDESOneGo function
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

// TestPackDESOne function
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

// TestTripleDESEncryptGo function
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

// TestTripleDESEncrypt function
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

// TestDESEncryptGo function
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

// TestDESEncrypt function
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

// BenchmarkPack3DES function
func BenchmarkPack3DES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_3des.txt"
		err := Pack3DES(src, dest)
		if err != nil {
			b.Fatal("Error Pack DES:", err)
		}
	}
}

// BenchmarkPackDES function
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

// BenchmarkPackDESWorkCalculate function
func BenchmarkPackDESWorkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		_, err := PackDESWorkCalculate(src)
		if err != nil {
			b.Fatal("Error Pack DES Work Calculate:", err)
		}
	}
}

// BenchmarkPack3DESOneGo function
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

// BenchmarkPack3DESOne function
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

// BenchmarkPackDESOneGo function
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

// BenchmarkPackDESOne function
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

// BenchmarkTripleDESEncryptGo function
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

// BenchmarkTripleDESEncrypt function
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

// BenchmarkDESEncryptGo function
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

// BenchmarkDESEncrypt function
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
