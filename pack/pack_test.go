package pack

import "testing"

func TestPack(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "AES"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

func TestPack2(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "DES"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

func TestPack3(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "3DES"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

func TestPack4(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "RSA"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

func TestPack5(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "BASE64"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

func BenchmarkPack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_aes.txt"
		algorithm := "AES"
		err := Pack(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Pack:", err)
		}
	}
}

func BenchmarkPack2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_aes.txt"
		algorithm := "DES"
		err := Pack(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Pack:", err)
		}
	}
}

func BenchmarkPack3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_aes.txt"
		algorithm := "3DES"
		err := Pack(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Pack:", err)
		}
	}
}

func BenchmarkPack4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_aes.txt"
		algorithm := "RSA"
		err := Pack(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Pack:", err)
		}
	}
}

func BenchmarkPack5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_aes.txt"
		algorithm := "BASE64"
		err := Pack(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error Pack:", err)
		}
	}
}
