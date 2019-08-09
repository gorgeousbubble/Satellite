package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var Help bool

func init() {
	flag.BoolVar(&Help, "h", false, "show help")

	flag.Usage = usage
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `Satellite Project - We share beautiful dreams together.
Version: v1.00a
Author: alopex

Usage: satellite [-h]
Options:
`)
	if err != nil {
		log.Println("Error print information:", err)
		os.Exit(1)
	}
	flag.PrintDefaults()
}
