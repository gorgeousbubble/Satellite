package cmd

import (
	"flag"
	. "satellite/global"
)

var qrCmd = flag.NewFlagSet(CmdQRCode, flag.ExitOnError)
var qrContent string
var qrSize int
var qrDest string

func init() {
	qrCmd.StringVar(&qrContent, "c", "https://github.com/gorgeousbubble", "content: the content string you want to show in qrcode, like url, for example: https://github.com/gorgeousbubble")
	qrCmd.IntVar(&qrSize, "s", 256, "size: the size of qrcode images width or length, like 128, 256, 512")
	qrCmd.StringVar(&qrDest, "o", "", "output files: one file path of qrcode iamges which you want to save, such as \"../test/data/iamges\"")
}
