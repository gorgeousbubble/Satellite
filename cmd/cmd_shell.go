package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/shell"
)

var shellCmd = flag.NewFlagSet(CmdShell, flag.ExitOnError)
var shellSrc string
var shellDest string
var shellType string

func init() {
	shellCmd.StringVar(&shellSrc, "i", "", "input file: the executable file which will be shelled")
	shellCmd.StringVar(&shellDest, "o", "", "output file: the output path after shell")
	shellCmd.StringVar(&shellType, "t", "upx", "shell type: choose one shell type from enum [upx]")
}

// ParseCmdShell function
// this function will be called in main.go and parse and execute one cmd command
// src is a binary or exe file which will be shelled by shell algorithm, like upx
// dest is a new file store path after binary or exe file shelled by shell algorithm
// algorithm is shell algorithm enum list, now support ['upx']
// any failure or error will output print to screen and exit process
func ParseCmdShell() {
	// check args number
	if len(os.Args) == 2 {
		shellCmd.Usage()
		os.Exit(1)
	}
	// parse command shell
	err := shellCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Shell Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdShell(shellSrc, shellDest, shellType)
	if err != nil {
		fmt.Print("\n")
		fmt.Println("Shell failure:", err)
		os.Exit(1)
	}
	fmt.Print("\n")
	fmt.Println("Shell success.")
}

// handleCmdShell function
// this function mainly handle the main flow of command
// any error will break and exit
func handleCmdShell(src string, dest string, algorithm string) (err error) {
	// check parameters
	if src == "" {
		err = errors.New("src file can not be empty")
		return err
	}
	if dest == "" {
		err = errors.New("dest file can not be empty")
		return err
	}
	if algorithm != "upx" {
		err = errors.New("shell type should be one of enum [upx]")
		return err
	}
	// shell executable file
	r, err := shell.Shell(src, dest, algorithm)
	if err != nil {
		log.Println("Error shell executable file:", err)
		fmt.Println(r)
		return err
	}
	fmt.Println(r)
	return err
}
