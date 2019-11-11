package images

import (
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestQRCodeGenerateToFile(t *testing.T) {
	err := QRCodeGenerateToFile("https://localhost:8080/", qrcode.Highest, 256, "../test/data/images/qrcode/qr_satellite.png")
	if err != nil {
		t.Errorf("Error generate qrcode to file: %v", err)
	}
}

func BenchmarkQRCodeGenerateToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := QRCodeGenerateToFile("https://localhost:8080/", qrcode.Highest, 256, "../test/data/images/qrcode/qr_satellite.png")
		if err != nil {
			b.Errorf("Error generate qrcode to file: %v", err)
		}
	}
}

func TestQRCodeGenerateToMemory(t *testing.T) {
	_, err := QRCodeGenerateToMemory("https://localhost:8080/", qrcode.Highest, 256)
	if err != nil {
		t.Errorf("Error generate qrcode to memory: %v", err)
	}
}

func BenchmarkQRCodeGenerateToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := QRCodeGenerateToMemory("https://localhost:8080/", qrcode.Highest, 256)
		if err != nil {
			b.Errorf("Error generate qrcode to memory: %v", err)
		}
	}
}
