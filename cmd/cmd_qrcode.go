package cmd

import (
	"Satellite/images"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"

	"github.com/skip2/go-qrcode"
)

var qrCmd = flag.NewFlagSet(CmdQRCode, flag.ExitOnError)
var qrContent string
var qrLevel string
var qrSize int
var qrDest string

func init() {
	qrCmd.StringVar(&qrContent, "c", "https://github.com/gorgeousbubble", "content: the content string you want to show in qrcode, like url, for example: https://github.com/gorgeousbubble")
	qrCmd.StringVar(&qrLevel, "l", "highest", "level: the images quality, you can choose one from 'low' 'medium' 'high' 'highest'")
	qrCmd.IntVar(&qrSize, "s", 256, "size: the size of qrcode images width or length, like 128, 256, 512")
	qrCmd.StringVar(&qrDest, "o", "", "output files: one file path of qrcode iamges which you want to save, such as \"../test/data/iamges/qr_satellite.png\"")
}

func ParseCmdQRCode() {
	// check args number
	if len(os.Args) == 2 {
		qrCmd.Usage()
		os.Exit(1)
	}
	// parse command qrcode
	err := qrCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse QRCode Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdQRCode(qrContent, qrLevel, qrSize, qrDest)
	if err != nil {
		fmt.Print("\n")
		fmt.Println("QRCode Generate failure:", err)
		os.Exit(1)
	}
	fmt.Print("\n")
	fmt.Println("QRCode Generate success.")
}

func handleCmdQRCode(content string, level string, size int, dest string) (err error) {
	var qrlevel qrcode.RecoveryLevel
	// check parameters
	if content == "" {
		err = errors.New("content can't be empty")
		return err
	}
	switch level {
	case "low":
		qrlevel = qrcode.Low
	case "medium":
		qrlevel = qrcode.Medium
	case "high":
		qrlevel = qrcode.High
	case "highest":
		qrlevel = qrcode.Highest
	default:
		err = errors.New("level should be one of 'low' 'medium' 'high' 'highest'")
		return err
	}
	if size <= 0 {
		err = errors.New("size shout more than zero")
		return err
	}
	// Generate QRCode function
	err = images.QRCodeGenerateToFile(content, qrlevel, size, dest)
	if err != nil {
		log.Println("Error generate QRCode:", err)
		return err
	}
	return err
}
