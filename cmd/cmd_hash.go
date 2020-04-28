package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
)

var hashCmd = flag.NewFlagSet(CmdHash, flag.ExitOnError)
var hashSrc string
var hashType string

func init() {
	hashCmd.StringVar(&hashSrc, "i", "", "hash source: one file or string which want to be calc by hash, such as \"test.txt\" or \"hello,world!\"")
	hashCmd.StringVar(&hashType, "t", "", "hash type: one type of enum [md5, sha1, sha256, sha512]")
}

func ParseCmdHash() {
	// check args number
	if len(os.Args) == 2 {
		hashCmd.Usage()
		os.Exit(1)
	}
	// parse command hash
	err := hashCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Hash Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdHash(hashSrc, hashType)
	if err != nil {
		fmt.Print("\n")
		fmt.Println("Hash calculate failure:", err)
		os.Exit(1)
	}
	fmt.Print("\n")
	fmt.Println("Hash calculate success.")
}

func handleCmdHash(src string, algorithm string) (err error) {
	// check parameters
	return err
}
