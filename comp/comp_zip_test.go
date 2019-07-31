package comp

import "testing"

func TestCompressZipOne(t *testing.T) {
	src := "../test/data/comp/file.txt"
	dest := "../test/data/comp/file.zip"
	err := CompressZipOne(src, dest)
	if err != nil {
		t.Fatal("Error Compress Zip One:", err)
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
