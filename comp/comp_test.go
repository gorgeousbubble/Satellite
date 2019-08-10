package comp

import "testing"

func TestCompress(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.tar.gz"
	algorithm := "tar"

	err := Compress(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Compress:", err)
	}
}

func TestCompress2(t *testing.T) {
	src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
	dest := "../test/data/comp/file.zip"
	algorithm := "zip"

	err := Compress(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Compress:", err)
	}
}

func BenchmarkCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/comp/file_5.txt"}
		dest := "../test/data/comp/file.tar.gz"
		algorithm := "tar"

		err := Compress(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Compress:", err)
		}
	}
}

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
