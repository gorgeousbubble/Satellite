package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"satellite/comp"
	. "satellite/global"
	"satellite/pack"
)

var compCmd = flag.NewFlagSet(CmdCompress, flag.ExitOnError)
var compSrc []string
var compDest string
var compType string

func init() {
	compCmd.Var(NewStrSlice([]string{}, &compSrc), "i", "input files: file list to compress, such as \"file_1.txt,file_2.mov,file_3.png...\"")
	compCmd.StringVar(&compDest, "o", "", "output files: one file end with 'tar.gz' or 'zip', such as \"file.tar.gz\" or \"file.zip\"")
	compCmd.StringVar(&compType, "t", "zip", "pack type: one type of enum [tar.gz,zip]")
}

func ParseCmdComp() {
	// check args number
	if len(os.Args) == 2 {
		compCmd.Usage()
		os.Exit(1)
	}
	// parse command comp
	err := compCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Comp Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdComp(compSrc, compDest, compType)
	if err != nil {
		fmt.Println("Compress failure:", err)
		os.Exit(1)
	}
	fmt.Println("Compress success.")
}

func handleCmdComp(src []string, dest string, algorithm string) (err error) {
	// execute pack function
	err = comp.Compress(src, dest, algorithm)
	if err != nil {
		log.Println("Compress failure:", err)
		return err
	}
	log.Println("Compress success.")
	return err
}
