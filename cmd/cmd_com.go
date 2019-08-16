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

Usage: satellite [pack/unpack] [comp/decomp] [tcp/udp]

Options:
	pack 	- packet files to user customize type.
	unpack	- unpack packet to origin files.
	comp	- compress files to 'zip' or 'tar.gz'.
	decomp	- decompress packet to origin files.
	tcp	- tcp simple server/client.
	udp	- udp simple server/client.
`)
	if err != nil {
		log.Println("Error print information:", err)
		os.Exit(1)
	}
	flag.PrintDefaults()
}
