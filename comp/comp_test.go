package comp

import "testing"

// TestCompress function
func TestCompress(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.tar.gz"
	algorithm := "tar.gz"

	err := Compress(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Compress:", err)
	}
}

// TestCompress2 function
func TestCompress2(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.zip"
	algorithm := "zip"

	err := Compress(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Compress:", err)
	}
}

// TestCompress3 function
func TestCompress3(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.tar"
	algorithm := "tar"

	err := Compress(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Compress:", err)
	}
}

// BenchmarkCompress function
func BenchmarkCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
		dest := "../test/data/comp/file.tar.gz"
		algorithm := "tar.gz"

		err := Compress(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Compress:", err)
		}
	}
}

// BenchmarkCompress2 function
func BenchmarkCompress2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
		dest := "../test/data/comp/file.zip"
		algorithm := "zip"

		err := Compress(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Compress:", err)
		}
	}
}

// BenchmarkCompress3 function
func BenchmarkCompress3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
		dest := "../test/data/comp/file.tar"
		algorithm := "tar"

		err := Compress(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Compress:", err)
		}
	}
}
