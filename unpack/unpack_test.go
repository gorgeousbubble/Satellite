package unpack

import "testing"

func TestUnpack(t *testing.T) {
	srcfile := "../test/data/unpack/file_aes.txt"
	destpath := "../test/data/unpack/"
	err := Unpack(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack2(t *testing.T) {
	srcfile := "../test/data/unpack/file_des.txt"
	destpath := "../test/data/unpack/"
	err := Unpack(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack3(t *testing.T) {
	srcfile := "../test/data/unpack/file_3des.txt"
	destpath := "../test/data/unpack/"
	err := Unpack(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack4(t *testing.T) {
	srcfile := "../test/data/unpack/file_rsa.txt"
	destpath := "../test/data/unpack/"
	err := Unpack(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack5(t *testing.T) {
	srcfile := "../test/data/unpack/file_base64.txt"
	destpath := "../test/data/unpack/"
	err := Unpack(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func BenchmarkUnpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_aes.txt"
		destpath := "../test/data/unpack/"
		err := Unpack(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpack2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_des.txt"
		destpath := "../test/data/unpack/"
		err := Unpack(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpack3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_3des.txt"
		destpath := "../test/data/unpack/"
		err := Unpack(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpack4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_rsa.txt"
		destpath := "../test/data/unpack/"
		err := Unpack(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpack5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_base64.txt"
		destpath := "../test/data/unpack/"
		err := Unpack(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}
