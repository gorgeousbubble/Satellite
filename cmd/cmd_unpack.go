package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/unpack"
)

var unpackCmd = flag.NewFlagSet(CmdUnpack, flag.ExitOnError)
var unpackSrc string
var unpackDest string

func init() {
	unpackCmd.StringVar(&unpackSrc, "i", "", "input files: packet file, such as \"file.dat\" or \"file.pak\"")
	unpackCmd.StringVar(&unpackDest, "o", "", "output files: one or more origin files. (should be path not file)")
}

func ParseCmdUnpack() {
	// check args number
	if len(os.Args) == 2 {
		unpackCmd.Usage()
		os.Exit(1)
	}
	// parse command unpack
	err := unpackCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Unpack Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdUnpack(unpackSrc, unpackDest)
	if err != nil {
		fmt.Println("Unpack failure:", err)
		os.Exit(1)
	}
	fmt.Println("Unpack success.")
}

func handleCmdUnpack(src string, dest string) (err error) {
	// execute unpack function
	err = unpack.Unpack(src, dest)
	if err != nil {
		log.Println("Unpack failure:", err)
		return err
	}
	log.Println("Unpack success.")
	return err
}
