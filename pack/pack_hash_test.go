package pack

import (
	"fmt"
	"testing"
)

func TestPackHashEncodeMD5(t *testing.T) {
	src := "hello,world!"
	algorithm := "md5"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash md5 encode:", dest)
}

func BenchmarkPackHashEncodeMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "md5"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeSHA1(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha1"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha1 encode:", dest)
}

func BenchmarkPackHashEncodeSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "sha1"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeSHA256(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha256"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha256 encode:", dest)
}

func BenchmarkPackHashEncodeSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "sha256"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeSHA512(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha512"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha512 encode:", dest)
}

func BenchmarkPackHashEncodeSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "sha512"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeBlake2b128(t *testing.T) {
	src := "hello,world!"
	algorithm := "blake2b128"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash blake2b encode:", dest)
}

func BenchmarkPackHashEncodeBlake2b128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "blake2b128"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeBlake2b256(t *testing.T) {
	src := "hello,world!"
	algorithm := "blake2b256"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash blake2b encode:", dest)
}

func BenchmarkPackHashEncodeBlake2b256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "blake2b256"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeBlake2b512(t *testing.T) {
	src := "hello,world!"
	algorithm := "blake2b512"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash blake2b encode:", dest)
}

func BenchmarkPackHashEncodeBlake2b512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "blake2b512"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashCheckMD5(t *testing.T) {
	src := "hello,world!"
	dest := "c0e84e870874dd37ed0d164c7986f03a"
	algorithm := "md5"
	result, err := PackHashCheck(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error pack hash check:", err)
	}
	if !result {
		t.Fatal("Error pack hash check result")
	}
}

func BenchmarkPackHashCheckMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "c0e84e870874dd37ed0d164c7986f03a"
		algorithm := "md5"
		result, err := PackHashCheck(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error pack hash check:", err)
		}
		if !result {
			b.Fatal("Error pack hash check result")
		}
	}
}

func TestPackHashCheckSHA1(t *testing.T) {
	src := "hello,world!"
	dest := "4518135c05e0706c0a34168996517bb3f28d94b5"
	algorithm := "sha1"
	result, err := PackHashCheck(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error pack hash check:", err)
	}
	if !result {
		t.Fatal("Error pack hash check result")
	}
}

func BenchmarkPackHashCheckSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "4518135c05e0706c0a34168996517bb3f28d94b5"
		algorithm := "sha1"
		result, err := PackHashCheck(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error pack hash check:", err)
		}
		if !result {
			b.Fatal("Error pack hash check result")
		}
	}
}

func TestPackHashCheckSHA256(t *testing.T) {
	src := "hello,world!"
	dest := "ec1e0bd875226943ad0e8877bdba4ca449c4cb8591a5363921c9f1ee20084c34"
	algorithm := "sha256"
	result, err := PackHashCheck(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error pack hash check:", err)
	}
	if !result {
		t.Fatal("Error pack hash check result")
	}
}

func BenchmarkPackHashCheckSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "ec1e0bd875226943ad0e8877bdba4ca449c4cb8591a5363921c9f1ee20084c34"
		algorithm := "sha256"
		result, err := PackHashCheck(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error pack hash check:", err)
		}
		if !result {
			b.Fatal("Error pack hash check result")
		}
	}
}

func TestPackHashCheckSHA512(t *testing.T) {
	src := "hello,world!"
	dest := "fa9edcfdaab7a4165f8d2593f04077d46aca957493e7e181ca479582d519a299d967305294d5d46fb5e0158240441b94cd96510c2311bdc86870e5ebf3efe60c"
	algorithm := "sha512"
	result, err := PackHashCheck(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error pack hash check:", err)
	}
	if !result {
		t.Fatal("Error pack hash check result")
	}
}

func BenchmarkPackHashCheckSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "fa9edcfdaab7a4165f8d2593f04077d46aca957493e7e181ca479582d519a299d967305294d5d46fb5e0158240441b94cd96510c2311bdc86870e5ebf3efe60c"
		algorithm := "sha512"
		result, err := PackHashCheck(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error pack hash check:", err)
		}
		if !result {
			b.Fatal("Error pack hash check result")
		}
	}
}

func TestPackHashCheckBlake2b128(t *testing.T) {
	src := "hello,world!"
	dest := "1748e3d0f53508245851db4571424eee"
	algorithm := "blake2b128"
	result, err := PackHashCheck(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error pack hash check:", err)
	}
	if !result {
		t.Fatal("Error pack hash check result")
	}
}

func BenchmarkPackHashCheckBlake2b128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "1748e3d0f53508245851db4571424eee"
		algorithm := "blake2b128"
		result, err := PackHashCheck(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error pack hash check:", err)
		}
		if !result {
			b.Fatal("Error pack hash check result")
		}
	}
}

func TestPackHashCheckBlake2b256(t *testing.T) {
	src := "hello,world!"
	dest := "8268578331a07de98347abd8cf11addf924a4ea0ac75e4aec1bf3fe6cb314553"
	algorithm := "blake2b256"
	result, err := PackHashCheck(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error pack hash check:", err)
	}
	if !result {
		t.Fatal("Error pack hash check result")
	}
}

func BenchmarkPackHashCheckBlake2b256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "8268578331a07de98347abd8cf11addf924a4ea0ac75e4aec1bf3fe6cb314553"
		algorithm := "blake2b256"
		result, err := PackHashCheck(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error pack hash check:", err)
		}
		if !result {
			b.Fatal("Error pack hash check result")
		}
	}
}

func TestPackHashCheckBlake2b512(t *testing.T) {
	src := "hello,world!"
	dest := "9fda5311802ea6b2e54fe5da18206584aed4a488cf8a06553e8b98f11f61d2b2115b9aba70c61d29bca31f444059cdba2f262a60358d23b0f661f75d5b91213f"
	algorithm := "blake2b512"
	result, err := PackHashCheck(src, dest, algorithm)
	if err != nil {
		t.Fatal("Error pack hash check:", err)
	}
	if !result {
		t.Fatal("Error pack hash check result")
	}
}

func BenchmarkPackHashCheckBlake2b512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "9fda5311802ea6b2e54fe5da18206584aed4a488cf8a06553e8b98f11f61d2b2115b9aba70c61d29bca31f444059cdba2f262a60358d23b0f661f75d5b91213f"
		algorithm := "blake2b512"
		result, err := PackHashCheck(src, dest, algorithm)
		if err != nil {
			b.Fatal("Error pack hash check:", err)
		}
		if !result {
			b.Fatal("Error pack hash check result")
		}
	}
}
