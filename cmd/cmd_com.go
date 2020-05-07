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
	_, err := fmt.Fprintf(os.Stderr, `
======================================================================================================================================================================================
==        <<<              <<<<<<<<       <<<<<<<<<<<<<<<<<   <<<<<<<<<<<<<<<<<   <<<                 <<<                 <<<<<<<<<<<<<<<<<   <<<<<<<<<<<<<<<<<   <<<<<<<<<<<<<<<<< ==
==      <<<  <<<          <<<    <<<      <<<    <<<    <<<   <<<           <<<   <<<                 <<<                 <<<    <<<    <<<   <<<    <<<    <<<   <<<           <<< ==
==    <<<      <<<       <<<      <<<     <<<    <<<    <<<   <<<                 <<<                 <<<                        <<<          <<<    <<<    <<<   <<<               ==
==  <<<          <<<    <<<        <<<           <<<          <<<                 <<<                 <<<                        <<<                 <<<          <<<               ==
==  <<<          <<<   <<<          <<<          <<<          <<<                 <<<                 <<<                        <<<                 <<<          <<<               ==
==    <<<              <<<          <<<          <<<          <<<                 <<<                 <<<                        <<<                 <<<          <<<               ==
==      <<<            <<<          <<<          <<<          <<<           <<<   <<<                 <<<                        <<<                 <<<          <<<           <<< ==
==        <<<          <<<<<<<<<<<<<<<<          <<<          <<<<<<<<<<<<<<<<<   <<<                 <<<                        <<<                 <<<          <<<<<<<<<<<<<<<<< ==
==          <<<        <<<          <<<          <<<          <<<           <<<   <<<                 <<<                        <<<                 <<<          <<<           <<< ==
==            <<<      <<<          <<<          <<<          <<<                 <<<                 <<<                        <<<                 <<<          <<<               ==
== <<<          <<<    <<<          <<<          <<<          <<<                 <<<                 <<<                        <<<                 <<<          <<<               ==
== <<<          <<<    <<<          <<<          <<<          <<<                 <<<                 <<<                        <<<                 <<<          <<<               ==
==   <<<      <<<      <<<          <<<          <<<          <<<                 <<<                 <<<                        <<<                 <<<          <<<               ==
==     <<<  <<<        <<<          <<<          <<<          <<<           <<<   <<<           <<<   <<<           <<<   <<<    <<<    <<<          <<<          <<<           <<< ==
==        <<<          <<<          <<<          <<<          <<<<<<<<<<<<<<<<<   <<<<<<<<<<<<<<<<<   <<<<<<<<<<<<<<<<<   <<<<<<<<<<<<<<<<<          <<<          <<<<<<<<<<<<<<<<< ==
======================================================================================================================================================================================
Satellite Project - We share beautiful dreams together.
Brief: Satellite is a tool for both multiple command and web service develop by Go-lang.
Version: v1.00a
Author: alopex

Usage: satellite [help] [pack/unpack] [comp/decomp] [tcp/udp] [http/https/ftp/rpc] [qrcode] [shell] [parses]

Options:
	help	- help information about satellite application.
	pack 	- packet files to user customize type.
	unpack	- unpack packet to origin files.
	comp	- compress files to 'zip' or 'tar.gz'.
	decomp	- decompress packet to origin files.
	hash    - calculate hash code of string or file.
	tcp     - tcp simple server/client.
	udp     - udp simple server/client.
	http 	- http restful server.
	https 	- https restful server.
	ftp     - ftp server/client.
	rpc     - rpc server for GoApi support tcp and http.
	qrcode  - generate qrcode save as images.
	shell   - shell executable file.
	parses	- multiple file parser.
`)
	if err != nil {
		log.Println("Error print information:", err)
		os.Exit(1)
	}
	flag.PrintDefaults()
}
