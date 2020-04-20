package decomp

import "testing"

// TestDeCompress function
func TestDeCompress(t *testing.T) {
	src := "../test/data/decomp/file.tar.gz"
	dest := "../test/data/decomp/"
	err := DeCompress(src, dest, "tar.gz")
	if err != nil {
		t.Fatal("Error DeCompress:", err)
	}
}

// TestDeCompress2 function
func TestDeCompress2(t *testing.T) {
	src := "../test/data/decomp/file.zip"
	dest := "../test/data/decomp/"
	err := DeCompress(src, dest, "zip")
	if err != nil {
		t.Fatal("Error DeCompress:", err)
	}
}

// TestDeCompress3 function
func TestDeCompress3(t *testing.T) {
	src := "../test/data/decomp/file.tar"
	dest := "../test/data/decomp/"
	err := DeCompress(src, dest, "tar")
	if err != nil {
		t.Fatal("Error DeCompress:", err)
	}
}

func BenchmarkDeCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/decomp/file.tar.gz"
		dest := "../test/data/decomp/"
		err := DeCompress(src, dest, "tar.gz")
		if err != nil {
			b.Fatal("Error DeCompress:", err)
		}
	}
}

func BenchmarkDeCompress2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/decomp/file.zip"
		dest := "../test/data/decomp/"
		err := DeCompress(src, dest, "zip")
		if err != nil {
			b.Fatal("Error DeCompress:", err)
		}
	}
}

func BenchmarkDeCompress3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/decomp/file.tar"
		dest := "../test/data/decomp/"
		err := DeCompress(src, dest, "tar")
		if err != nil {
			b.Fatal("Error DeCompress:", err)
		}
	}
}
