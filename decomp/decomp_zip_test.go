package decomp

import "testing"

func TestDeCompressZip(t *testing.T) {
	src := "../test/data/decomp/file.zip"
	dest := "../test/data/decomp/"
	err := DeCompressZip(src, dest)
	if err != nil {
		t.Fatal("Error DeCompress Zip:", err)
	}
}

func BenchmarkDeCompressZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/decomp/file.zip"
		dest := "../test/data/decomp/"
		err := DeCompressZip(src, dest)
		if err != nil {
			b.Fatal("Error DeCompress Zip:", err)
		}
	}
}
