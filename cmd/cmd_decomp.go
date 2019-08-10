package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"satellite/decomp"
	. "satellite/global"
)

var deCompCmd = flag.NewFlagSet(CmdDecompress, flag.ExitOnError)
var deCompSrc string
var deCompDest string
var deCompType string

func init() {
	deCompCmd.StringVar(&deCompSrc, "i", "", "input files: compress file, such as \"file.tar.gz\" or \"file.zip\"")
	deCompCmd.StringVar(&deCompDest, "o", "", "output files: one or more origin files. (should be path not file)")
	deCompCmd.StringVar(&deCompType, "t", "zip", "decompress type: one type of enum [tar.gz,zip]")
}

func ParseCmdDeComp() {
	// check args number
	if len(os.Args) == 2 {
		deCompCmd.Usage()
		os.Exit(1)
	}
	// parse command unpack
	err := deCompCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse DeComp Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdDeComp(deCompSrc, deCompDest, deCompType)
	if err != nil {
		fmt.Println("Decompress failure:", err)
		os.Exit(1)
	}
	fmt.Println("Decompress success.")
}

func handleCmdDeComp(src string, dest string, algorithm string) (err error) {
	// execute unpack function
	err = decomp.DeCompress(src, dest, algorithm)
	if err != nil {
		log.Println("Decompress failure:", err)
		return err
	}
	log.Println("Decompress success.")
	return err
}
