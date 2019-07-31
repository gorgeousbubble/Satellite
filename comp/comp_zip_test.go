package comp

import (
	"testing"
)

func TestCompressZip(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.zip"
	err := CompressZip(src, dest)
	if err != nil {
		t.Fatal("Error Compress Zip:", err)
	}
}

func TestCompressZipOne(t *testing.T) {
	src := "../test/data/comp/file.txt"
	dest := "../test/data/comp/file.zip"
	err := CompressZipOne(src, dest)
	if err != nil {
		t.Fatal("Error Compress Zip One:", err)
	}
}

func BenchmarkCompressZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
		dest := "../test/data/comp/file.zip"
		err := CompressZip(src, dest)
		if err != nil {
			b.Fatal("Error Compress Zip:", err)
		}
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
