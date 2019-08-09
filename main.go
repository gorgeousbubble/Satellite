package main

import (
	"flag"
	"os"
	_ "satellite/cmd"
)

func main() {
	if len(os.Args) <= 1 {
		flag.Usage()
	}

}
