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
	hashCmd.StringVar(&hashType, "t", "", "hash type: one type of enum [md5, sha1, sha256, sha512, blake2b128, blake2b256, blake2b512, hmac_sha1, hmac_sha256, hmac_sha512]")
}

// ParseCmdHash function
// this function will be called in main.go and parse and execute one hash command
// input src string support sentences and file path, such as 'hello,world!' or '../data/hash/file.txt'
// algorithm now support 'md5', 'sha1', 'sha256', 'sha512', you can send both up case and low case
// any failure or error will output print to screen and exit process
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

// handleCmdHash function
// this function mainly handle the main flow of command
// any error will break and exit
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

// checkHashParameters function
// this function will check parameters legally
// return true indicate check pass, otherwise check failed
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
	case "BLAKE2B128", "blake2b128":
	case "BLAKE2B256", "blake2b256":
	case "BLAKE2B512", "blake2b512":
	case "HMAC_SHA1", "hmac_sha1":
	case "HMAC_SHA256", "hmac_sha256":
	case "HMAC_SHA512", "hmac_sha512":
	default:
		is = false
		fmt.Printf("Algorithm %v not support.\n", algorithm)
	}
	return is
}
