package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"satellite/app/erika/cli"
	_ "satellite/app/erika/logs"
)

func init() {
	// start multiple CPU core
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
}

func main() {
	// check cli args number
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	// switch cli execute
	switch os.Args[1] {
	case "help":
		flag.Usage()
	case "start":
		cli.ParseCliStart()
	case "stop":
		cli.ParseCliStop()
	default:
		fmt.Println("Unrecognized command~")
		os.Exit(1)
	}
}
