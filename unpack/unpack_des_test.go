package unpack

import (
	"bytes"
	"io/ioutil"
	. "satellite/utils"
	"sync"
	"testing"
)

func TestUnpackDES(t *testing.T) {
	srcfile := "../test/data/unpack/file_des.txt"
	destpath := "../test/data/unpack/"
	err := UnpackDES(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack DES:", err)
	}
}

func TestUnpackDESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xB8, 0xB5, 0xBA, 0x6E, 0x15, 0x0C, 0x90, 0xE4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
		0x28, 0xD4, 0x07, 0x26, 0xCB, 0x76, 0x16, 0x67, 0xB7, 0xA9, 0x6A, 0x69, 0x4E, 0x41, 0xD9, 0xD6,
		0x35, 0xB3, 0x3B, 0x2D, 0x7F, 0xCC, 0x4F, 0xBC, 0x33, 0x4C, 0x8B, 0x4F, 0xA6, 0xA2, 0x33, 0x2B,
		0xD3, 0x62, 0xDB, 0x19, 0xAA, 0x17, 0xAE, 0x7D, 0x60, 0xCD, 0x31, 0x40, 0x51, 0x1D, 0x9E, 0x65,
		0x68, 0x3E, 0xCD, 0xBB, 0x70, 0x2D, 0x8B, 0x0C, 0x71, 0x82, 0x46, 0x68, 0xAA, 0x5A, 0xB7, 0x41,
		0xE9, 0x76, 0x2D, 0x55, 0x30, 0xF3, 0x80, 0x94, 0xE7, 0x76, 0xDA, 0x0A, 0x40, 0x1B, 0xCC, 0xD3,
		0x7F, 0xB2, 0x5E, 0xAD, 0xAB, 0xDB, 0x07, 0xF3, 0x66, 0x47, 0xE5, 0x7C, 0xA2, 0xE5, 0x5B, 0x03,
		0xC7, 0x15, 0xF5, 0xCD, 0x8D, 0xD9, 0x92, 0xF6, 0xB2, 0xF0, 0x9D, 0xF0, 0x6A, 0xE3, 0xBE, 0x3B,
		0x1B, 0x9C, 0xAC, 0xBF, 0x9E, 0xBF, 0xA9, 0xE9, 0x0B, 0xED, 0xD0, 0xCB, 0x23, 0x62, 0x93, 0x91}
	destpath := "../test/data/unpack/"
	hh := TUnpackDESOne{}
	hh.Name = make([]byte, 32)
	hh.Key = make([]byte, 8)
	hh.OriginSize = make([]byte, 4)
	hh.CryptSize = make([]byte, 4)

	rd := bytes.NewReader(src)
	_, err := rd.Read(hh.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(hh.Key)
	if err != nil {
		t.Fatal("Error read header key:", err)
	}
	_, err = rd.Read(hh.OriginSize)
	if err != nil {
		t.Fatal("Error read header origin size:", err)
	}
	_, err = rd.Read(hh.CryptSize)
	if err != nil {
		t.Fatal("Error read header crypt size:", err)
	}
	s := make([]byte, BytesToInt(hh.CryptSize))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}

	wg.Add(1)
	go UnpackDESOneGo(s, hh, destpath, &wg)
	wg.Wait()
}

func TestUnpackDESOne(t *testing.T) {
	src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xB8, 0xB5, 0xBA, 0x6E, 0x15, 0x0C, 0x90, 0xE4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
		0x28, 0xD4, 0x07, 0x26, 0xCB, 0x76, 0x16, 0x67, 0xB7, 0xA9, 0x6A, 0x69, 0x4E, 0x41, 0xD9, 0xD6,
		0x35, 0xB3, 0x3B, 0x2D, 0x7F, 0xCC, 0x4F, 0xBC, 0x33, 0x4C, 0x8B, 0x4F, 0xA6, 0xA2, 0x33, 0x2B,
		0xD3, 0x62, 0xDB, 0x19, 0xAA, 0x17, 0xAE, 0x7D, 0x60, 0xCD, 0x31, 0x40, 0x51, 0x1D, 0x9E, 0x65,
		0x68, 0x3E, 0xCD, 0xBB, 0x70, 0x2D, 0x8B, 0x0C, 0x71, 0x82, 0x46, 0x68, 0xAA, 0x5A, 0xB7, 0x41,
		0xE9, 0x76, 0x2D, 0x55, 0x30, 0xF3, 0x80, 0x94, 0xE7, 0x76, 0xDA, 0x0A, 0x40, 0x1B, 0xCC, 0xD3,
		0x7F, 0xB2, 0x5E, 0xAD, 0xAB, 0xDB, 0x07, 0xF3, 0x66, 0x47, 0xE5, 0x7C, 0xA2, 0xE5, 0x5B, 0x03,
		0xC7, 0x15, 0xF5, 0xCD, 0x8D, 0xD9, 0x92, 0xF6, 0xB2, 0xF0, 0x9D, 0xF0, 0x6A, 0xE3, 0xBE, 0x3B,
		0x1B, 0x9C, 0xAC, 0xBF, 0x9E, 0xBF, 0xA9, 0xE9, 0x0B, 0xED, 0xD0, 0xCB, 0x23, 0x62, 0x93, 0x91}
	destpath := "../test/data/unpack/"
	hh := TUnpackDESOne{}
	hh.Name = make([]byte, 32)
	hh.Key = make([]byte, 8)
	hh.OriginSize = make([]byte, 4)
	hh.CryptSize = make([]byte, 4)

	rd := bytes.NewReader(src)
	_, err := rd.Read(hh.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(hh.Key)
	if err != nil {
		t.Fatal("Error read header key:", err)
	}
	_, err = rd.Read(hh.OriginSize)
	if err != nil {
		t.Fatal("Error read header origin size:", err)
	}
	_, err = rd.Read(hh.CryptSize)
	if err != nil {
		t.Fatal("Error read header crypt size:", err)
	}
	s := make([]byte, BytesToInt(hh.CryptSize))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}

	err = UnpackDESOne(s, hh, destpath)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
}

func TestDESDecryptGo(t *testing.T) {
	var dest []byte
	var wg sync.WaitGroup
	src := []byte{0x8F, 0x34, 0x1C, 0x84, 0xBB, 0xB4, 0x58, 0xA8, 0x54, 0x8B, 0xC0, 0x9E, 0x2F, 0x55, 0x01, 0xF2}
	key := []byte("hyacinth")
	wg.Add(1)
	go DESDecryptGo(src, key, &dest, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/unpack/file.txt", dest, 0644)
	if err != nil {
		t.Fatal("Error Write DES One:", err)
	}
}

func TestDESDecrypt(t *testing.T) {
	src := []byte{0x8F, 0x34, 0x1C, 0x84, 0xBB, 0xB4, 0x58, 0xA8, 0x54, 0x8B, 0xC0, 0x9E, 0x2F, 0x55, 0x01, 0xF2}
	key := []byte("hyacinth")
	r, err := DESDecrypt(src, key)
	if err != nil {
		t.Fatal("Error DES Decrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES One:", err)
	}
}

func BenchmarkUnpackDES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_des.txt"
		destpath := "../test/data/unpack/"
		err := UnpackDES(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack DES:", err)
		}
	}
}

func BenchmarkUnpackDESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0xB8, 0xB5, 0xBA, 0x6E, 0x15, 0x0C, 0x90, 0xE4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
			0x28, 0xD4, 0x07, 0x26, 0xCB, 0x76, 0x16, 0x67, 0xB7, 0xA9, 0x6A, 0x69, 0x4E, 0x41, 0xD9, 0xD6,
			0x35, 0xB3, 0x3B, 0x2D, 0x7F, 0xCC, 0x4F, 0xBC, 0x33, 0x4C, 0x8B, 0x4F, 0xA6, 0xA2, 0x33, 0x2B,
			0xD3, 0x62, 0xDB, 0x19, 0xAA, 0x17, 0xAE, 0x7D, 0x60, 0xCD, 0x31, 0x40, 0x51, 0x1D, 0x9E, 0x65,
			0x68, 0x3E, 0xCD, 0xBB, 0x70, 0x2D, 0x8B, 0x0C, 0x71, 0x82, 0x46, 0x68, 0xAA, 0x5A, 0xB7, 0x41,
			0xE9, 0x76, 0x2D, 0x55, 0x30, 0xF3, 0x80, 0x94, 0xE7, 0x76, 0xDA, 0x0A, 0x40, 0x1B, 0xCC, 0xD3,
			0x7F, 0xB2, 0x5E, 0xAD, 0xAB, 0xDB, 0x07, 0xF3, 0x66, 0x47, 0xE5, 0x7C, 0xA2, 0xE5, 0x5B, 0x03,
			0xC7, 0x15, 0xF5, 0xCD, 0x8D, 0xD9, 0x92, 0xF6, 0xB2, 0xF0, 0x9D, 0xF0, 0x6A, 0xE3, 0xBE, 0x3B,
			0x1B, 0x9C, 0xAC, 0xBF, 0x9E, 0xBF, 0xA9, 0xE9, 0x0B, 0xED, 0xD0, 0xCB, 0x23, 0x62, 0x93, 0x91}
		destpath := "../test/data/unpack/"
		hh := TUnpackDESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 8)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)

		rd := bytes.NewReader(src)
		_, err := rd.Read(hh.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			b.Fatal("Error read header key:", err)
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			b.Fatal("Error read header origin size:", err)
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			b.Fatal("Error read header crypt size:", err)
		}
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}

		wg.Add(1)
		go UnpackDESOneGo(s, hh, destpath, &wg)
		wg.Wait()
	}
}

func BenchmarkUnpackDESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0xB8, 0xB5, 0xBA, 0x6E, 0x15, 0x0C, 0x90, 0xE4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
			0x28, 0xD4, 0x07, 0x26, 0xCB, 0x76, 0x16, 0x67, 0xB7, 0xA9, 0x6A, 0x69, 0x4E, 0x41, 0xD9, 0xD6,
			0x35, 0xB3, 0x3B, 0x2D, 0x7F, 0xCC, 0x4F, 0xBC, 0x33, 0x4C, 0x8B, 0x4F, 0xA6, 0xA2, 0x33, 0x2B,
			0xD3, 0x62, 0xDB, 0x19, 0xAA, 0x17, 0xAE, 0x7D, 0x60, 0xCD, 0x31, 0x40, 0x51, 0x1D, 0x9E, 0x65,
			0x68, 0x3E, 0xCD, 0xBB, 0x70, 0x2D, 0x8B, 0x0C, 0x71, 0x82, 0x46, 0x68, 0xAA, 0x5A, 0xB7, 0x41,
			0xE9, 0x76, 0x2D, 0x55, 0x30, 0xF3, 0x80, 0x94, 0xE7, 0x76, 0xDA, 0x0A, 0x40, 0x1B, 0xCC, 0xD3,
			0x7F, 0xB2, 0x5E, 0xAD, 0xAB, 0xDB, 0x07, 0xF3, 0x66, 0x47, 0xE5, 0x7C, 0xA2, 0xE5, 0x5B, 0x03,
			0xC7, 0x15, 0xF5, 0xCD, 0x8D, 0xD9, 0x92, 0xF6, 0xB2, 0xF0, 0x9D, 0xF0, 0x6A, 0xE3, 0xBE, 0x3B,
			0x1B, 0x9C, 0xAC, 0xBF, 0x9E, 0xBF, 0xA9, 0xE9, 0x0B, 0xED, 0xD0, 0xCB, 0x23, 0x62, 0x93, 0x91}
		destpath := "../test/data/unpack/"
		hh := TUnpackDESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 8)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)

		rd := bytes.NewReader(src)
		_, err := rd.Read(hh.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			b.Fatal("Error read header key:", err)
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			b.Fatal("Error read header origin size:", err)
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			b.Fatal("Error read header crypt size:", err)
		}
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}

		err = UnpackDESOne(s, hh, destpath)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
	}
}

func BenchmarkDESDecryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		var wg sync.WaitGroup
		src := []byte{0x8F, 0x34, 0x1C, 0x84, 0xBB, 0xB4, 0x58, 0xA8, 0x54, 0x8B, 0xC0, 0x9E, 0x2F, 0x55, 0x01, 0xF2}
		key := []byte("hyacinth")
		wg.Add(1)
		go DESDecryptGo(src, key, &dest, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/unpack/file.txt", dest, 0644)
		if err != nil {
			b.Fatal("Error Write DES One:", err)
		}
	}
}

func BenchmarkDESDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{0x8F, 0x34, 0x1C, 0x84, 0xBB, 0xB4, 0x58, 0xA8, 0x54, 0x8B, 0xC0, 0x9E, 0x2F, 0x55, 0x01, 0xF2}
		key := []byte("hyacinth")
		r, err := DESDecrypt(src, key)
		if err != nil {
			b.Fatal("Error DES Decrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES One:", err)
		}
	}
}
