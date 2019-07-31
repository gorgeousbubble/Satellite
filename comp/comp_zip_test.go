package comp

import (
	"sync"
	"testing"
)

func TestCompressZipOneGo(t *testing.T) {
	var wg sync.WaitGroup
	src := "../test/data/comp/file.txt"
	dest := "../test/data/comp/file.zip"
	wg.Add(1)
	go CompressZipOneGo(src, dest, &wg)
	wg.Wait()
}

func TestCompressZipOne(t *testing.T) {
	src := "../test/data/comp/file.txt"
	dest := "../test/data/comp/file.zip"
	err := CompressZipOne(src, dest)
	if err != nil {
		t.Fatal("Error Compress Zip One:", err)
	}
}

func BenchmarkCompressZipOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		src := "../test/data/comp/file.txt"
		dest := "../test/data/comp/file.zip"
		wg.Add(1)
		go CompressZipOneGo(src, dest, &wg)
		wg.Wait()
	}
}

func BenchmarkCompressZipOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/comp/file.txt"
		dest := "../test/data/comp/file.zip"
		err := CompressZipOne(src, dest)
		if err != nil {
			b.Fatal("Error Compress Zip One:", err)
		}
	}
}
