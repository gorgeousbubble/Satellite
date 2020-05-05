package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/pack"
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
	is := checkHashParameters(src, algorithm)
	if !is {
		err = errors.New("parameters illegal")
		return err
	}
	// calculate hash
	dest, err := pack.PackHashEncode(src, algorithm)
	if err != nil {
		return err
	}
	fmt.Printf("Hash %v:%v\n", algorithm, dest)
	return err
}

func checkHashParameters(src string, algorithm string) (is bool) {
	// check input source
	is = true
	// check src
	if len(src) == 0 {
		is = false
		fmt.Println("Source string can't be empty.")
		return is
	}
	// check algorithm
	switch algorithm {
	case "MD5", "md5":
	case "SHA1", "sha1":
	case "SHA256", "sha256":
	case "SHA512", "sha512":
	default:
		is = false
		fmt.Printf("Algorithm %v not support.\n", algorithm)
	}
	return is
}
