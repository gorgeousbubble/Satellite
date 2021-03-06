package cli

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
-------------------------------------------------------------------------------------------------------------------------
--  +++++++++++++++++   +++++++++++++++++         +++++++++++++++++   +++++          ++++++            +++++           --
--  +++           +++   +++            +++        +++    +++    +++    +++            +++             +++++++          --
--  +++                 +++              +++             +++           +++          +++              +++   +++         --
--  +++                 +++              +++             +++           +++        +++                +++   +++         --
--  +++                 +++              +++             +++           +++      +++                 +++     +++        --
--  +++                 +++              +++             +++           +++    +++                   +++     +++        --
--  +++           +++   +++            +++               +++           +++  +++                    +++       +++       --
--  +++++++++++++++++   +++++++++++++++++                +++           +++++++                     +++++++++++++       --
--  +++           +++   +++           +++                +++           +++  +++                   +++         +++      --
--  +++                 +++            +++               +++           +++    +++                 +++         +++      --
--  +++                 +++             +++              +++           +++      +++              +++           +++     --
--  +++                 +++              +++             +++           +++        +++            +++           +++     --
--  +++                 +++               +++            +++           +++          +++         +++             +++    --
--  +++           +++   +++                +++    +++    +++    +++    +++            +++       +++             +++    --
--  +++++++++++++++++   +++                 +++   +++++++++++++++++   +++++          ++++++   ++++++           ++++++  --
-------------------------------------------------------------------------------------------------------------------------
Erika Project - We share beautiful dreams together.
Brief: Erika is a artificial intelligence(AI) develop by Go.
Version: v1.00a
Author: alopex

Usage: erika [help] [start/stop] [status]

Options:
    help    - help information about erika application.
    start   - start an erika instance.
    stop    - stop all erika instance.
`)
	if err != nil {
		log.Println("Error print information:", err)
		os.Exit(1)
	}
	flag.PrintDefaults()
}
