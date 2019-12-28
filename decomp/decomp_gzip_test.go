package decomp

import "testing"

func TestDeCompressGzip(t *testing.T) {
	src := []string{"../test/data/decomp/file_1.gz", "../test/data/decomp/file_2.gz", "../test/data/decomp/file_3.gz", "../test/data/decomp/file_4.gz", "../test/data/decomp/file_5.gz"}
	dest := "../test/data/decomp"
	err := DeCompressGzip(src, dest)
	if err != nil {
		t.Fatal("Error DeCompress gzip:", err)
	}
}
