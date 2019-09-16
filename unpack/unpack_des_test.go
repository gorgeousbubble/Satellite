package unpack

import (
	"bytes"
	"io/ioutil"
	. "satellite/utils"
	"sync"
	"testing"
)

func TestUnpack3DES(t *testing.T) {
	src := "../test/data/unpack/file_3des.txt"
	dest := "../test/data/unpack/"
	err := Unpack3DES(src, dest)
	if err != nil {
		t.Fatal("Error Unpack 3DES:", err)
	}
}

func TestUnpackDES(t *testing.T) {
	src := "../test/data/unpack/file_des.txt"
	dest := "../test/data/unpack/"
	err := UnpackDES(src, dest)
	if err != nil {
		t.Fatal("Error Unpack DES:", err)
	}
}

func TestUnpack3DESExtractInfo(t *testing.T) {
	var dest []string
	src := "../test/data/unpack/file_3des.txt"
	err := Unpack3DESExtractInfo(src, &dest)
	if err != nil {
		t.Fatal("Error Unpack 3DES Extract Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract file number")
	}
}

func TestUnpackDESExtractInfo(t *testing.T) {
	var dest []string
	src := "../test/data/unpack/file_des.txt"
	err := UnpackDESExtractInfo(src, &dest)
	if err != nil {
		t.Fatal("Error Unpack DES Extract Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract file number")
	}
}

func TestUnpack3DESWorkCalculate(t *testing.T) {
	src := "../test/data/unpack/file_3des.txt"
	_, err := Unpack3DESWorkCalculate(src)
	if err != nil {
		t.Fatal("Error Unpack 3DES Work Calculate:", err)
	}
}

func TestUnpackDESWorkCalculate(t *testing.T) {
	src := "../test/data/unpack/file_des.txt"
	_, err := UnpackDESWorkCalculate(src)
	if err != nil {
		t.Fatal("Error Unpack DES Work Calculate:", err)
	}
}

func TestUnpack3DESOneToMemory(t *testing.T) {
	var dest []byte
	src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x02, 0x28, 0xB6, 0x44, 0x10, 0xF4, 0x1A, 0x4A, 0x8E, 0x4B, 0xE9, 0x53, 0xB2, 0xD9, 0x7D, 0xC6,
		0x51, 0x87, 0xED, 0xEC, 0x99, 0xB9, 0xC9, 0xB4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
		0x27, 0x98, 0x0E, 0xE9, 0xD2, 0xA0, 0x5D, 0x29, 0x81, 0xE0, 0x43, 0xB7, 0xB3, 0x26, 0x4D, 0xD0,
		0x36, 0xA9, 0xFF, 0x15, 0xC9, 0x12, 0x82, 0xD3, 0xA4, 0x1B, 0x73, 0x8D, 0x6E, 0xF9, 0x2D, 0x2F,
		0xF4, 0x93, 0x13, 0xE2, 0xC1, 0x08, 0x60, 0x21, 0x15, 0x55, 0xD9, 0x22, 0xDD, 0x20, 0xE5, 0xAF,
		0xE2, 0x69, 0xFA, 0x61, 0x93, 0x0B, 0x90, 0x9D, 0x9D, 0x6A, 0xE4, 0xA6, 0xE4, 0xA1, 0x83, 0x9A,
		0x15, 0xD4, 0x7D, 0x55, 0x3D, 0xE4, 0x52, 0x7F, 0x8F, 0x37, 0x37, 0x74, 0xF7, 0x92, 0x7C, 0x29,
		0x9D, 0x40, 0x48, 0x9C, 0xB0, 0xBA, 0x2B, 0xED, 0xD3, 0xEF, 0x50, 0x82, 0xAE, 0x4C, 0x2F, 0x4C,
		0x03, 0x43, 0x37, 0xBD, 0x4E, 0x8D, 0x9B, 0xE4, 0x03, 0xC2, 0x6D, 0xE6, 0x23, 0x0C, 0x81, 0x71,
		0x8D, 0xA4, 0x6A, 0x29, 0x05, 0x35, 0x19, 0xA1, 0x0E, 0x55, 0x39, 0xF5, 0x8C, 0x62, 0x73, 0x3A}
	hh := TUnpack3DESOne{}
	hh.Name = make([]byte, 32)
	hh.Key = make([]byte, 24)
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

	err = Unpack3DESOneToMemory(s, hh, &dest)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
	if !bytes.Equal(dest, []byte("hello,world!")) {
		t.Fatal("Error unpack content:", string(dest))
	}
}

func TestUnpack3DESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x02, 0x28, 0xB6, 0x44, 0x10, 0xF4, 0x1A, 0x4A, 0x8E, 0x4B, 0xE9, 0x53, 0xB2, 0xD9, 0x7D, 0xC6,
		0x51, 0x87, 0xED, 0xEC, 0x99, 0xB9, 0xC9, 0xB4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
		0x27, 0x98, 0x0E, 0xE9, 0xD2, 0xA0, 0x5D, 0x29, 0x81, 0xE0, 0x43, 0xB7, 0xB3, 0x26, 0x4D, 0xD0,
		0x36, 0xA9, 0xFF, 0x15, 0xC9, 0x12, 0x82, 0xD3, 0xA4, 0x1B, 0x73, 0x8D, 0x6E, 0xF9, 0x2D, 0x2F,
		0xF4, 0x93, 0x13, 0xE2, 0xC1, 0x08, 0x60, 0x21, 0x15, 0x55, 0xD9, 0x22, 0xDD, 0x20, 0xE5, 0xAF,
		0xE2, 0x69, 0xFA, 0x61, 0x93, 0x0B, 0x90, 0x9D, 0x9D, 0x6A, 0xE4, 0xA6, 0xE4, 0xA1, 0x83, 0x9A,
		0x15, 0xD4, 0x7D, 0x55, 0x3D, 0xE4, 0x52, 0x7F, 0x8F, 0x37, 0x37, 0x74, 0xF7, 0x92, 0x7C, 0x29,
		0x9D, 0x40, 0x48, 0x9C, 0xB0, 0xBA, 0x2B, 0xED, 0xD3, 0xEF, 0x50, 0x82, 0xAE, 0x4C, 0x2F, 0x4C,
		0x03, 0x43, 0x37, 0xBD, 0x4E, 0x8D, 0x9B, 0xE4, 0x03, 0xC2, 0x6D, 0xE6, 0x23, 0x0C, 0x81, 0x71,
		0x8D, 0xA4, 0x6A, 0x29, 0x05, 0x35, 0x19, 0xA1, 0x0E, 0x55, 0x39, 0xF5, 0x8C, 0x62, 0x73, 0x3A}
	dest := "../test/data/unpack/"
	hh := TUnpack3DESOne{}
	hh.Name = make([]byte, 32)
	hh.Key = make([]byte, 24)
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
	go Unpack3DESOneGo(s, hh, dest, &wg)
	wg.Wait()
}

func TestUnpack3DESOne(t *testing.T) {
	src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x02, 0x28, 0xB6, 0x44, 0x10, 0xF4, 0x1A, 0x4A, 0x8E, 0x4B, 0xE9, 0x53, 0xB2, 0xD9, 0x7D, 0xC6,
		0x51, 0x87, 0xED, 0xEC, 0x99, 0xB9, 0xC9, 0xB4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
		0x27, 0x98, 0x0E, 0xE9, 0xD2, 0xA0, 0x5D, 0x29, 0x81, 0xE0, 0x43, 0xB7, 0xB3, 0x26, 0x4D, 0xD0,
		0x36, 0xA9, 0xFF, 0x15, 0xC9, 0x12, 0x82, 0xD3, 0xA4, 0x1B, 0x73, 0x8D, 0x6E, 0xF9, 0x2D, 0x2F,
		0xF4, 0x93, 0x13, 0xE2, 0xC1, 0x08, 0x60, 0x21, 0x15, 0x55, 0xD9, 0x22, 0xDD, 0x20, 0xE5, 0xAF,
		0xE2, 0x69, 0xFA, 0x61, 0x93, 0x0B, 0x90, 0x9D, 0x9D, 0x6A, 0xE4, 0xA6, 0xE4, 0xA1, 0x83, 0x9A,
		0x15, 0xD4, 0x7D, 0x55, 0x3D, 0xE4, 0x52, 0x7F, 0x8F, 0x37, 0x37, 0x74, 0xF7, 0x92, 0x7C, 0x29,
		0x9D, 0x40, 0x48, 0x9C, 0xB0, 0xBA, 0x2B, 0xED, 0xD3, 0xEF, 0x50, 0x82, 0xAE, 0x4C, 0x2F, 0x4C,
		0x03, 0x43, 0x37, 0xBD, 0x4E, 0x8D, 0x9B, 0xE4, 0x03, 0xC2, 0x6D, 0xE6, 0x23, 0x0C, 0x81, 0x71,
		0x8D, 0xA4, 0x6A, 0x29, 0x05, 0x35, 0x19, 0xA1, 0x0E, 0x55, 0x39, 0xF5, 0x8C, 0x62, 0x73, 0x3A}
	dest := "../test/data/unpack/"
	hh := TUnpack3DESOne{}
	hh.Name = make([]byte, 32)
	hh.Key = make([]byte, 24)
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

	err = Unpack3DESOne(s, hh, dest)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
}

func TestUnpackDESOneToMemory(t *testing.T) {
	var dest []byte
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

	err = UnpackDESOneToMemory(s, hh, &dest)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
	if !bytes.Equal(dest, []byte("hello,world!")) {
		t.Fatal("Error unpack content:", string(dest))
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
	dest := "../test/data/unpack/"
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
	go UnpackDESOneGo(s, hh, dest, &wg)
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
	dest := "../test/data/unpack/"
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

	err = UnpackDESOne(s, hh, dest)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
}

func TestTripleDESDecryptGo(t *testing.T) {
	var dest []byte
	var wg sync.WaitGroup
	src := []byte{0x2F, 0x58, 0x70, 0x9E, 0xA3, 0x98, 0x7E, 0xD7, 0x8E, 0x79, 0x12, 0x0F, 0x47, 0xF3, 0xA5, 0x49}
	key := []byte("HyacinthRaindropRomantic")
	wg.Add(1)
	go TripleDESDecryptGo(src, key, &dest, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/unpack/file.txt", dest, 0644)
	if err != nil {
		t.Fatal("Error Write 3DES One:", err)
	}
}

func TestTripleDESDecrypt(t *testing.T) {
	src := []byte{0x2F, 0x58, 0x70, 0x9E, 0xA3, 0x98, 0x7E, 0xD7, 0x8E, 0x79, 0x12, 0x0F, 0x47, 0xF3, 0xA5, 0x49}
	key := []byte("HyacinthRaindropRomantic")
	r, err := TripleDESDecrypt(src, key)
	if err != nil {
		t.Fatal("Error 3DES Decrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write 3DES One:", err)
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

func BenchmarkUnpack3DES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_3des.txt"
		dest := "../test/data/unpack/"
		err := Unpack3DES(src, dest)
		if err != nil {
			b.Fatal("Error Unpack 3DES:", err)
		}
	}
}

func BenchmarkUnpackDES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_des.txt"
		dest := "../test/data/unpack/"
		err := UnpackDES(src, dest)
		if err != nil {
			b.Fatal("Error Unpack DES:", err)
		}
	}
}

func BenchmarkUnpack3DESExtractInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		src := "../test/data/unpack/file_3des.txt"
		err := Unpack3DESExtractInfo(src, &dest)
		if err != nil {
			b.Fatal("Error Unpack 3DES Extract Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract file number")
		}
	}
}

func BenchmarkUnpackDESExtractInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		src := "../test/data/unpack/file_des.txt"
		err := UnpackDESExtractInfo(src, &dest)
		if err != nil {
			b.Fatal("Error Unpack DES Extract Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract file number")
		}
	}
}

func BenchmarkUnpack3DESWorkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_3des.txt"
		_, err := Unpack3DESWorkCalculate(src)
		if err != nil {
			b.Fatal("Error Unpack 3DES Work Calculate:", err)
		}
	}
}

func BenchmarkUnpackDESWorkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_des.txt"
		_, err := UnpackDESWorkCalculate(src)
		if err != nil {
			b.Fatal("Error Unpack DES Work Calculate:", err)
		}
	}
}

func BenchmarkUnpack3DESOneToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x02, 0x28, 0xB6, 0x44, 0x10, 0xF4, 0x1A, 0x4A, 0x8E, 0x4B, 0xE9, 0x53, 0xB2, 0xD9, 0x7D, 0xC6,
			0x51, 0x87, 0xED, 0xEC, 0x99, 0xB9, 0xC9, 0xB4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
			0x27, 0x98, 0x0E, 0xE9, 0xD2, 0xA0, 0x5D, 0x29, 0x81, 0xE0, 0x43, 0xB7, 0xB3, 0x26, 0x4D, 0xD0,
			0x36, 0xA9, 0xFF, 0x15, 0xC9, 0x12, 0x82, 0xD3, 0xA4, 0x1B, 0x73, 0x8D, 0x6E, 0xF9, 0x2D, 0x2F,
			0xF4, 0x93, 0x13, 0xE2, 0xC1, 0x08, 0x60, 0x21, 0x15, 0x55, 0xD9, 0x22, 0xDD, 0x20, 0xE5, 0xAF,
			0xE2, 0x69, 0xFA, 0x61, 0x93, 0x0B, 0x90, 0x9D, 0x9D, 0x6A, 0xE4, 0xA6, 0xE4, 0xA1, 0x83, 0x9A,
			0x15, 0xD4, 0x7D, 0x55, 0x3D, 0xE4, 0x52, 0x7F, 0x8F, 0x37, 0x37, 0x74, 0xF7, 0x92, 0x7C, 0x29,
			0x9D, 0x40, 0x48, 0x9C, 0xB0, 0xBA, 0x2B, 0xED, 0xD3, 0xEF, 0x50, 0x82, 0xAE, 0x4C, 0x2F, 0x4C,
			0x03, 0x43, 0x37, 0xBD, 0x4E, 0x8D, 0x9B, 0xE4, 0x03, 0xC2, 0x6D, 0xE6, 0x23, 0x0C, 0x81, 0x71,
			0x8D, 0xA4, 0x6A, 0x29, 0x05, 0x35, 0x19, 0xA1, 0x0E, 0x55, 0x39, 0xF5, 0x8C, 0x62, 0x73, 0x3A}
		hh := TUnpack3DESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 24)
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

		err = Unpack3DESOneToMemory(s, hh, &dest)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
		if !bytes.Equal(dest, []byte("hello,world!")) {
			b.Fatal("Error unpack content:", string(dest))
		}
	}
}

func BenchmarkUnpack3DESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x02, 0x28, 0xB6, 0x44, 0x10, 0xF4, 0x1A, 0x4A, 0x8E, 0x4B, 0xE9, 0x53, 0xB2, 0xD9, 0x7D, 0xC6,
			0x51, 0x87, 0xED, 0xEC, 0x99, 0xB9, 0xC9, 0xB4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
			0x27, 0x98, 0x0E, 0xE9, 0xD2, 0xA0, 0x5D, 0x29, 0x81, 0xE0, 0x43, 0xB7, 0xB3, 0x26, 0x4D, 0xD0,
			0x36, 0xA9, 0xFF, 0x15, 0xC9, 0x12, 0x82, 0xD3, 0xA4, 0x1B, 0x73, 0x8D, 0x6E, 0xF9, 0x2D, 0x2F,
			0xF4, 0x93, 0x13, 0xE2, 0xC1, 0x08, 0x60, 0x21, 0x15, 0x55, 0xD9, 0x22, 0xDD, 0x20, 0xE5, 0xAF,
			0xE2, 0x69, 0xFA, 0x61, 0x93, 0x0B, 0x90, 0x9D, 0x9D, 0x6A, 0xE4, 0xA6, 0xE4, 0xA1, 0x83, 0x9A,
			0x15, 0xD4, 0x7D, 0x55, 0x3D, 0xE4, 0x52, 0x7F, 0x8F, 0x37, 0x37, 0x74, 0xF7, 0x92, 0x7C, 0x29,
			0x9D, 0x40, 0x48, 0x9C, 0xB0, 0xBA, 0x2B, 0xED, 0xD3, 0xEF, 0x50, 0x82, 0xAE, 0x4C, 0x2F, 0x4C,
			0x03, 0x43, 0x37, 0xBD, 0x4E, 0x8D, 0x9B, 0xE4, 0x03, 0xC2, 0x6D, 0xE6, 0x23, 0x0C, 0x81, 0x71,
			0x8D, 0xA4, 0x6A, 0x29, 0x05, 0x35, 0x19, 0xA1, 0x0E, 0x55, 0x39, 0xF5, 0x8C, 0x62, 0x73, 0x3A}
		dest := "../test/data/unpack/"
		hh := TUnpack3DESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 24)
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
		go Unpack3DESOneGo(s, hh, dest, &wg)
		wg.Wait()
	}
}

func BenchmarkUnpack3DESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x02, 0x28, 0xB6, 0x44, 0x10, 0xF4, 0x1A, 0x4A, 0x8E, 0x4B, 0xE9, 0x53, 0xB2, 0xD9, 0x7D, 0xC6,
			0x51, 0x87, 0xED, 0xEC, 0x99, 0xB9, 0xC9, 0xB4, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80,
			0x27, 0x98, 0x0E, 0xE9, 0xD2, 0xA0, 0x5D, 0x29, 0x81, 0xE0, 0x43, 0xB7, 0xB3, 0x26, 0x4D, 0xD0,
			0x36, 0xA9, 0xFF, 0x15, 0xC9, 0x12, 0x82, 0xD3, 0xA4, 0x1B, 0x73, 0x8D, 0x6E, 0xF9, 0x2D, 0x2F,
			0xF4, 0x93, 0x13, 0xE2, 0xC1, 0x08, 0x60, 0x21, 0x15, 0x55, 0xD9, 0x22, 0xDD, 0x20, 0xE5, 0xAF,
			0xE2, 0x69, 0xFA, 0x61, 0x93, 0x0B, 0x90, 0x9D, 0x9D, 0x6A, 0xE4, 0xA6, 0xE4, 0xA1, 0x83, 0x9A,
			0x15, 0xD4, 0x7D, 0x55, 0x3D, 0xE4, 0x52, 0x7F, 0x8F, 0x37, 0x37, 0x74, 0xF7, 0x92, 0x7C, 0x29,
			0x9D, 0x40, 0x48, 0x9C, 0xB0, 0xBA, 0x2B, 0xED, 0xD3, 0xEF, 0x50, 0x82, 0xAE, 0x4C, 0x2F, 0x4C,
			0x03, 0x43, 0x37, 0xBD, 0x4E, 0x8D, 0x9B, 0xE4, 0x03, 0xC2, 0x6D, 0xE6, 0x23, 0x0C, 0x81, 0x71,
			0x8D, 0xA4, 0x6A, 0x29, 0x05, 0x35, 0x19, 0xA1, 0x0E, 0x55, 0x39, 0xF5, 0x8C, 0x62, 0x73, 0x3A}
		dest := "../test/data/unpack/"
		hh := TUnpack3DESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 24)
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

		err = Unpack3DESOne(s, hh, dest)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
	}
}

func BenchmarkUnpackDESOneToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
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

		err = UnpackDESOneToMemory(s, hh, &dest)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
		if !bytes.Equal(dest, []byte("hello,world!")) {
			b.Fatal("Error unpack content:", string(dest))
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
		dest := "../test/data/unpack/"
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
		go UnpackDESOneGo(s, hh, dest, &wg)
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
		dest := "../test/data/unpack/"
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

		err = UnpackDESOne(s, hh, dest)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
	}
}

func BenchmarkTripleDESDecryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		var wg sync.WaitGroup
		src := []byte{0x2F, 0x58, 0x70, 0x9E, 0xA3, 0x98, 0x7E, 0xD7, 0x8E, 0x79, 0x12, 0x0F, 0x47, 0xF3, 0xA5, 0x49}
		key := []byte("HyacinthRaindropRomantic")
		wg.Add(1)
		go TripleDESDecryptGo(src, key, &dest, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/unpack/file.txt", dest, 0644)
		if err != nil {
			b.Fatal("Error Write 3DES One:", err)
		}
	}
}

func BenchmarkTripleDESDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{0x2F, 0x58, 0x70, 0x9E, 0xA3, 0x98, 0x7E, 0xD7, 0x8E, 0x79, 0x12, 0x0F, 0x47, 0xF3, 0xA5, 0x49}
		key := []byte("HyacinthRaindropRomantic")
		r, err := TripleDESDecrypt(src, key)
		if err != nil {
			b.Fatal("Error 3DES Decrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write 3DES One:", err)
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
