package unpack

import "testing"

// TestUnpack function
func TestUnpack(t *testing.T) {
	src := "../test/data/unpack/file_aes.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpack2 function
func TestUnpack2(t *testing.T) {
	src := "../test/data/unpack/file_des.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpack3 function
func TestUnpack3(t *testing.T) {
	src := "../test/data/unpack/file_3des.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpack4 function
func TestUnpack4(t *testing.T) {
	src := "../test/data/unpack/file_rsa.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpack5 function
func TestUnpack5(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpackConfine function
func TestUnpackConfine(t *testing.T) {
	src := "../test/data/unpack/file_aes.txt"
	dest := "../test/data/unpack/"
	err := UnpackConfine(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpackConfine2 function
func TestUnpackConfine2(t *testing.T) {
	src := "../test/data/unpack/file_des.txt"
	dest := "../test/data/unpack/"
	err := UnpackConfine(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpackConfine3 function
func TestUnpackConfine3(t *testing.T) {
	src := "../test/data/unpack/file_3des.txt"
	dest := "../test/data/unpack/"
	err := UnpackConfine(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpackConfine4 function
func TestUnpackConfine4(t *testing.T) {
	src := "../test/data/unpack/file_rsa.txt"
	dest := "../test/data/unpack/"
	err := UnpackConfine(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpackConfine5 function
func TestUnpackConfine5(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	err := UnpackConfine(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

// TestUnpackToFile function
func TestUnpackToFile(t *testing.T) {
	src := "../test/data/unpack/file_aes.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFile(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFile2 function
func TestUnpackToFile2(t *testing.T) {
	src := "../test/data/unpack/file_des.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFile(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFile3 function
func TestUnpackToFile3(t *testing.T) {
	src := "../test/data/unpack/file_3des.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFile(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFile4 function
func TestUnpackToFile4(t *testing.T) {
	src := "../test/data/unpack/file_rsa.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFile(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFile5 function
func TestUnpackToFile5(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFile(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFileConfine function
func TestUnpackToFileConfine(t *testing.T) {
	src := "../test/data/unpack/file_aes.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFileConfine(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFileConfine2 function
func TestUnpackToFileConfine2(t *testing.T) {
	src := "../test/data/unpack/file_des.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFileConfine(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFileConfine3 function
func TestUnpackToFileConfine3(t *testing.T) {
	src := "../test/data/unpack/file_3des.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFileConfine(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFileConfine4 function
func TestUnpackToFileConfine4(t *testing.T) {
	src := "../test/data/unpack/file_rsa.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFileConfine(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToFileConfine5 function
func TestUnpackToFileConfine5(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackToFileConfine(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack To File:", err)
	}
}

// TestUnpackToMemory function
func TestUnpackToMemory(t *testing.T) {
	var dest []byte
	src := "../test/data/unpack/file_aes.txt"
	target := "file_1.txt"
	err := UnpackToMemory(src, target, &dest)
	if err != nil {
		t.Fatal("Error Unpack To Memory:", err)
	}
}

// TestUnpackToMemory2 function
func TestUnpackToMemory2(t *testing.T) {
	var dest []byte
	src := "../test/data/unpack/file_des.txt"
	target := "file_1.txt"
	err := UnpackToMemory(src, target, &dest)
	if err != nil {
		t.Fatal("Error Unpack To Memory:", err)
	}
}

// TestUnpackToMemory3 function
func TestUnpackToMemory3(t *testing.T) {
	var dest []byte
	src := "../test/data/unpack/file_3des.txt"
	target := "file_1.txt"
	err := UnpackToMemory(src, target, &dest)
	if err != nil {
		t.Fatal("Error Unpack To Memory:", err)
	}
}

// TestUnpackToMemory4 function
func TestUnpackToMemory4(t *testing.T) {
	var dest []byte
	src := "../test/data/unpack/file_rsa.txt"
	target := "file_1.txt"
	err := UnpackToMemory(src, target, &dest)
	if err != nil {
		t.Fatal("Error Unpack To Memory:", err)
	}
}

// TestUnpackToMemory5 function
func TestUnpackToMemory5(t *testing.T) {
	var dest []byte
	src := "../test/data/unpack/file_base64.txt"
	target := "file_1.txt"
	err := UnpackToMemory(src, target, &dest)
	if err != nil {
		t.Fatal("Error Unpack To Memory:", err)
	}
}

// TestExtractInfo function
func TestExtractInfo(t *testing.T) {
	var dest []string
	var size []int
	var algorithm string
	src := "../test/data/unpack/file_aes.txt"
	err := ExtractInfo(src, &dest, &size, &algorithm)
	if err != nil {
		t.Fatal("Error Extract AES Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract Number")
	}
}

// TestExtractInfo2 function
func TestExtractInfo2(t *testing.T) {
	var dest []string
	var size []int
	var algorithm string
	src := "../test/data/unpack/file_base64.txt"
	err := ExtractInfo(src, &dest, &size, &algorithm)
	if err != nil {
		t.Fatal("Error Extract BASE64 Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract Number")
	}
}

// TestExtractInfo3 function
func TestExtractInfo3(t *testing.T) {
	var dest []string
	var size []int
	var algorithm string
	src := "../test/data/unpack/file_3des.txt"
	err := ExtractInfo(src, &dest, &size, &algorithm)
	if err != nil {
		t.Fatal("Error Extract 3DES Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract Number")
	}
}

// TestExtractInfo4 function
func TestExtractInfo4(t *testing.T) {
	var dest []string
	var size []int
	var algorithm string
	src := "../test/data/unpack/file_des.txt"
	err := ExtractInfo(src, &dest, &size, &algorithm)
	if err != nil {
		t.Fatal("Error Extract DES Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract Number")
	}
}

// TestExtractInfo5 function
func TestExtractInfo5(t *testing.T) {
	var dest []string
	var size []int
	var algorithm string
	src := "../test/data/unpack/file_rsa.txt"
	err := ExtractInfo(src, &dest, &size, &algorithm)
	if err != nil {
		t.Fatal("Error Extract RSA Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract Number")
	}
}

// TestWorkCalculate function
func TestWorkCalculate(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_aes.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate2 function
func TestWorkCalculate2(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_des.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate3 function
func TestWorkCalculate3(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_3des.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate4 function
func TestWorkCalculate4(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_rsa.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// TestWorkCalculate5 function
func TestWorkCalculate5(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_base64.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

// BenchmarkUnpack function
func BenchmarkUnpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_aes.txt"
		dest := "../test/data/unpack/"
		err := Unpack(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

// BenchmarkUnpack2 function
func BenchmarkUnpack2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_des.txt"
		dest := "../test/data/unpack/"
		err := Unpack(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

// BenchmarkUnpack3 function
func BenchmarkUnpack3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_3des.txt"
		dest := "../test/data/unpack/"
		err := Unpack(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

// BenchmarkUnpack4 function
func BenchmarkUnpack4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_rsa.txt"
		dest := "../test/data/unpack/"
		err := Unpack(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

// BenchmarkUnpack5 function
func BenchmarkUnpack5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		err := Unpack(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpackConfine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_aes.txt"
		dest := "../test/data/unpack/"
		err := UnpackConfine(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpackConfine2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_des.txt"
		dest := "../test/data/unpack/"
		err := UnpackConfine(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpackConfine3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_3des.txt"
		dest := "../test/data/unpack/"
		err := UnpackConfine(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpackConfine4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_rsa.txt"
		dest := "../test/data/unpack/"
		err := UnpackConfine(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpackConfine5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		err := UnpackConfine(src, dest)
		if err != nil {
			b.Fatal("Error Unpack:", err)
		}
	}
}

func BenchmarkUnpackToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_aes.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFile(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFile2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_des.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFile(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFile3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_3des.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFile(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFile4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_rsa.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFile(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFile5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFile(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFileConfine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_aes.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFileConfine(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFileConfine2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_des.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFileConfine(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFileConfine3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_3des.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFileConfine(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFileConfine4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_rsa.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFileConfine(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToFileConfine5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackToFileConfine(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack To File:", err)
		}
	}
}

func BenchmarkUnpackToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		src := "../test/data/unpack/file_aes.txt"
		target := "file_1.txt"
		err := UnpackToMemory(src, target, &dest)
		if err != nil {
			b.Fatal("Error Unpack To Memory:", err)
		}
	}
}

func BenchmarkUnpackToMemory2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		src := "../test/data/unpack/file_des.txt"
		target := "file_1.txt"
		err := UnpackToMemory(src, target, &dest)
		if err != nil {
			b.Fatal("Error Unpack To Memory:", err)
		}
	}
}

func BenchmarkUnpackToMemory3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		src := "../test/data/unpack/file_3des.txt"
		target := "file_1.txt"
		err := UnpackToMemory(src, target, &dest)
		if err != nil {
			b.Fatal("Error Unpack To Memory:", err)
		}
	}
}

func BenchmarkUnpackToMemory4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		src := "../test/data/unpack/file_rsa.txt"
		target := "file_1.txt"
		err := UnpackToMemory(src, target, &dest)
		if err != nil {
			b.Fatal("Error Unpack To Memory:", err)
		}
	}
}

func BenchmarkUnpackToMemory5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		src := "../test/data/unpack/file_base64.txt"
		target := "file_1.txt"
		err := UnpackToMemory(src, target, &dest)
		if err != nil {
			b.Fatal("Error Unpack To Memory:", err)
		}
	}
}

func BenchmarkExtractInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		var size []int
		var algorithm string
		src := "../test/data/unpack/file_aes.txt"
		err := ExtractInfo(src, &dest, &size, &algorithm)
		if err != nil {
			b.Fatal("Error Extract AES Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract Number")
		}
	}
}

func BenchmarkExtractInfo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		var size []int
		var algorithm string
		src := "../test/data/unpack/file_base64.txt"
		err := ExtractInfo(src, &dest, &size, &algorithm)
		if err != nil {
			b.Fatal("Error Extract BASE64 Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract Number")
		}
	}
}

func BenchmarkExtractInfo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		var size []int
		var algorithm string
		src := "../test/data/unpack/file_3des.txt"
		err := ExtractInfo(src, &dest, &size, &algorithm)
		if err != nil {
			b.Fatal("Error Extract 3DES Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract Number")
		}
	}
}

func BenchmarkExtractInfo4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		var size []int
		var algorithm string
		src := "../test/data/unpack/file_des.txt"
		err := ExtractInfo(src, &dest, &size, &algorithm)
		if err != nil {
			b.Fatal("Error Extract DES Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract Number")
		}
	}
}

func BenchmarkExtractInfo5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		var size []int
		var algorithm string
		src := "../test/data/unpack/file_rsa.txt"
		err := ExtractInfo(src, &dest, &size, &algorithm)
		if err != nil {
			b.Fatal("Error Extract RSA Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract Number")
		}
	}
}

func BenchmarkWorkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		var algorithm string
		src := "../test/data/unpack/file_aes.txt"
		err := WorkCalculate(src, &algorithm, &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		var algorithm string
		src := "../test/data/unpack/file_des.txt"
		err := WorkCalculate(src, &algorithm, &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		var algorithm string
		src := "../test/data/unpack/file_3des.txt"
		err := WorkCalculate(src, &algorithm, &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		var algorithm string
		src := "../test/data/unpack/file_rsa.txt"
		err := WorkCalculate(src, &algorithm, &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}

func BenchmarkWorkCalculate5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var work int64
		var algorithm string
		src := "../test/data/unpack/file_base64.txt"
		err := WorkCalculate(src, &algorithm, &work)
		if err != nil {
			b.Fatal("Error Work Calculate:", err)
		}
	}
}
