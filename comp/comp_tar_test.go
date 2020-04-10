package comp

import "testing"

// TestCompressTar function
func TestCompressTar(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.tar"
	err := CompressTar(src, dest)
	if err != nil {
		t.Fatal("Error Compress Tar:", err)
	}
}

// BenchmarkCompressTar function
func BenchmarkCompressTar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
		dest := "../test/data/comp/file.tar"
		err := CompressTar(src, dest)
		if err != nil {
			b.Fatal("Error Compress Tar:", err)
		}
	}
}

// TestCompressTarGz function
func TestCompressTarGz(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.tar.gz"
	err := CompressTarGz(src, dest)
	if err != nil {
		t.Fatal("Error Compress Tar Gz:", err)
	}
}

// BenchmarkCompressTarGz function
func BenchmarkCompressTarGz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
		dest := "../test/data/comp/file.tar.gz"
		err := CompressTarGz(src, dest)
		if err != nil {
			b.Fatal("Error Compress Tar Gz:", err)
		}
	}
}
