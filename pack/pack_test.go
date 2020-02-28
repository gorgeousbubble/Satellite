package pack

import "testing"

// TestPack function
func TestPack(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "AES"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

// TestPack2 function
func TestPack2(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "DES"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

// TestPack3 function
func TestPack3(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "3DES"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

// TestPack4 function
func TestPack4(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "RSA"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

// TestPack5 function
func TestPack5(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_aes.txt"
	algorithm := "BASE64"
	err := Pack(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error Pack:", err)
	}
}

// TestWorkCalculate function
func TestWorkCalculate(t *testing.T) {
	var work int64
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	err := WorkCalculate(src, "aes", &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate2 function
func TestWorkCalculate2(t *testing.T) {
	var work int64
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	err := WorkCalculate(src, "des", &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate3 function
func TestWorkCalculate3(t *testing.T) {
	var work int64
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	err := WorkCalculate(src, "3des", &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate4 function
func TestWorkCalculate4(t *testing.T) {
	var work int64
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	err := WorkCalculate(src, "rsa", &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate5 function
func TestWorkCalculate5(t *testing.T) {
	var work int64
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	err := WorkCalculate(src, "base64", &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
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

func BenchmarkWorkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		err := WorkCalculate(src, "aes", &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		err := WorkCalculate(src, "des", &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		err := WorkCalculate(src, "3des", &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		err := WorkCalculate(src, "rsa", &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		err := WorkCalculate(src, "base64", &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}
