package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
)

var parsesCmd = flag.NewFlagSet(CmdParses, flag.ExitOnError)
var parsesResolver string
var parsesSrc string
var parsesName string
var parsesType string

func init() {
	parsesCmd.StringVar(&parsesResolver, "r", "", "parses package resolver, you can choose one from [\"ini\",\"erl\"]")
	parsesCmd.StringVar(&parsesSrc, "i", "", "input file: correspond file need to be parse, such as \"test.ini\" or \"test.erl\"")
	parsesCmd.StringVar(&parsesName, "n", "", "parameter name: correspond parameter need to be parse, such as \"conf\" or \"pi\"")
	parsesCmd.StringVar(&parsesType, "t", "", "parameter type: correspond parameter type need to be parse, such as \"bool\" or \"int\"")
}

func ParseCmdParses() {
	// check args number
	if len(os.Args) == 2 {
		parsesCmd.Usage()
		os.Exit(1)
	}
	// parse command parses
	err := parsesCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Parses Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdParses(parsesResolver, parsesSrc, parsesName, parsesType)
	if err != nil {
		fmt.Println("Parses failure:", err)
		os.Exit(1)
	}
	fmt.Println("Parses success.")
}

func handleCmdParses(resolver string, src string, name string, kind string) (err error) {
	// handle command parses
	return err
}
