package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"satellite/cmd"
	. "satellite/global"
	_ "satellite/logging"
)

func init() {
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
}

func main() {
	// check command args number
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// switch command execute
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
