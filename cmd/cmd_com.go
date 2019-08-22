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

Usage: satellite [pack/unpack] [comp/decomp] [tcp/udp] [http/https]

Options:
	pack 	- packet files to user customize type.
	unpack	- unpack packet to origin files.
	comp	- compress files to 'zip' or 'tar.gz'.
	decomp	- decompress packet to origin files.
	tcp	- tcp simple server/client.
	udp	- udp simple server/client.
	http 	- http restful server
	https 	- https restful server
`)
	if err != nil {
		log.Println("Error print information:", err)
		os.Exit(1)
	}
	flag.PrintDefaults()
}
