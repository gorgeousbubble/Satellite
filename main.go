package main

import (
	"flag"
	"fmt"
	"os"
	"satellite/cmd"
	. "satellite/global"
)

func main() {
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case CmdPacket:
		cmd.ParseCmdPack()
	case CmdUnpack:
		cmd.ParseCmdUnpack()
	default:
		fmt.Println("Unexpected command~")
		os.Exit(1)
	}

}
