package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	flag.Usage = Usage
}

func Usage() {
	_, err := fmt.Fprintf(os.Stderr, `Satellite Project - We share beautiful dreams together.
Version: v1.00a
Author: alopex

Usage: satellite [pack/unpack]

Options:
`)
	if err != nil {
		log.Println("Error print information:", err)
		os.Exit(1)
	}
	flag.PrintDefaults()
}
