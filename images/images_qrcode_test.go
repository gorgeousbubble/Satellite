package images

import (
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestQRCodeGenerate(t *testing.T) {
	err := QRCodeGenerate("https://localhost:8080/", qrcode.Highest, 256, "../test/data/images/qrcode/qr_satellite.png")
	if err != nil {
		t.Errorf("Error generate qrcode: %v", err)
	}
}

func BenchmarkQRCodeGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := QRCodeGenerate("https://localhost:8080/", qrcode.Highest, 256, "../test/data/images/qrcode/qr_satellite.png")
		if err != nil {
			b.Errorf("Error generate qrcode: %v", err)
		}
	}
}
