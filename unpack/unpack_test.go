package unpack

import "testing"

func TestUnpack(t *testing.T) {
	src := "../test/data/unpack/file_aes.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack2(t *testing.T) {
	src := "../test/data/unpack/file_des.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack3(t *testing.T) {
	src := "../test/data/unpack/file_3des.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack4(t *testing.T) {
	src := "../test/data/unpack/file_rsa.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestUnpack5(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	err := Unpack(src, dest)
	if err != nil {
		t.Fatal("Error Unpack:", err)
	}
}

func TestWorkCalculate(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_aes.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

func TestWorkCalculate2(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_des.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

func TestWorkCalculate3(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_3des.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

func TestWorkCalculate4(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_rsa.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

func TestWorkCalculate5(t *testing.T) {
	var work int64
	var algorithm string
	src := "../test/data/unpack/file_base64.txt"
	err := WorkCalculate(src, &algorithm, &work)
	if err != nil {
		t.Fatal("Error Work Calculate:", err)
	}
}

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
