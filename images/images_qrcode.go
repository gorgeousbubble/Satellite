package images

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func QRCodeGenerate(content string, level qrcode.RecoveryLevel, size int, filename string) (err error) {
	err = qrcode.WriteFile(content, level, size, filename)
	if err != nil {
		log.Println("Error generate qr code:", err)
	}
	return err
}
