package decomp

import "testing"

// TestDeCompressGzip function
func TestDeCompressGzip(t *testing.T) {
	src := []string{"../test/data/decomp/file_1.gz", "../test/data/decomp/file_2.gz", "../test/data/decomp/file_3.gz", "../test/data/decomp/file_4.gz", "../test/data/decomp/file_5.gz"}
	dest := "../test/data/decomp"
	err := DeCompressGzip(src, dest)
	if err != nil {
		t.Fatal("Error DeCompress gzip:", err)
	}
}

// BenchmarkDeCompressGzip function
func BenchmarkDeCompressGzip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/decomp/file_1.gz", "../test/data/decomp/file_2.gz", "../test/data/decomp/file_3.gz", "../test/data/decomp/file_4.gz", "../test/data/decomp/file_5.gz"}
		dest := "../test/data/decomp"
		err := DeCompressGzip(src, dest)
		if err != nil {
			b.Fatal("Error DeCompress gzip:", err)
		}
	}
}
