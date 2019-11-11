package images

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func QRCodeGenerateToFile(content string, level qrcode.RecoveryLevel, size int, filename string) (err error) {
	err = qrcode.WriteFile(content, level, size, filename)
	if err != nil {
		log.Println("Error generate qr code:", err)
		return err
	}
	return err
}

func QRCodeGenerateToMemory(content string, level qrcode.RecoveryLevel, size int) (qr []byte, err error) {
	qr, err = qrcode.Encode(content, level, size)
	if err != nil {
		log.Println("Error generate qr code to memory", err)
		return qr, err
	}
	return qr, err
}
