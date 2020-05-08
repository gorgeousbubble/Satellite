package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"satellite/comp"
	. "satellite/global"
)

var compCmd = flag.NewFlagSet(CmdCompress, flag.ExitOnError)
var compSrc []string
var compDest string
var compType string

func init() {
	compCmd.Var(NewStrSlice([]string{}, &compSrc), "i", "input files: file list to compress, such as \"file_1.txt,file_2.mov,file_3.png...\"")
	compCmd.StringVar(&compDest, "o", "", "output files: one file end with 'tar.gz' or 'zip', such as \"file.tar.gz\" or \"file.zip\"")
	compCmd.StringVar(&compType, "t", "zip", "compress type: one type of enum [tar.gz,zip]")
}

// ParseCmdComp function
// this function will be called in main.go and parse and execute one comp command
// comp command has three parameters
// input src file list, output dest file path and algorithm which used in unpack, return error info
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// algorithm now support 'tar', 'tar.gz', 'zip', you can send both up case and low case
// any failure or error will output print to screen and exit process
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

// handleCmdComp function
// this function mainly handle the main flow of command
// any error will break and exit
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
