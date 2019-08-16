package main

import (
	"flag"
	"fmt"
	"os"
	"satellite/cmd"
	. "satellite/global"
	_ "satellite/logging"
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
	case CmdCompress:
		cmd.ParseCmdComp()
	case CmdDecompress:
		cmd.ParseCmdDeComp()
	case CmdTcp:
		cmd.ParseCmdTcp()
	case CmdUdp:
		cmd.ParseCmdUdp()
	default:
		fmt.Println("Invalid command~")
		os.Exit(1)
	}
}
