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
#########################################################################################################################
##  $$$$$$$$$$$$$$$$$   $$$$$$$$$$$$$$$$$         $$$$$$$$$$$$$$$$$   $$$$$          $$$$$$            $$$$$           ##
##  $$$           $$$   $$$            $$$        $$$    $$$    $$$    $$$            $$$             $$$$$$$          ##
##  $$$                 $$$              $$$             $$$           $$$          $$$              $$$   $$$         ##
##  $$$                 $$$              $$$             $$$           $$$        $$$                $$$   $$$         ##
##  $$$                 $$$              $$$             $$$           $$$      $$$                 $$$     $$$        ##
##  $$$                 $$$              $$$             $$$           $$$    $$$                   $$$     $$$        ##
##  $$$           $$$   $$$            $$$               $$$           $$$  $$$                    $$$       $$$       ##
##  $$$$$$$$$$$$$$$$$   $$$$$$$$$$$$$$$$$                $$$           $$$$$$$                     $$$$$$$$$$$$$       ##
##  $$$           $$$   $$$           $$$                $$$           $$$  $$$                   $$$         $$$      ##
##  $$$                 $$$            $$$               $$$           $$$    $$$                 $$$         $$$      ##
##  $$$                 $$$             $$$              $$$           $$$      $$$              $$$           $$$     ##
##  $$$                 $$$              $$$             $$$           $$$        $$$            $$$           $$$     ##
##  $$$                 $$$               $$$            $$$           $$$          $$$         $$$             $$$    ##
##  $$$           $$$   $$$                $$$    $$$    $$$    $$$    $$$            $$$       $$$             $$$    ##
##  $$$$$$$$$$$$$$$$$   $$$                 $$$   $$$$$$$$$$$$$$$$$   $$$$$          $$$$$$   $$$$$$           $$$$$$  ##
#########################################################################################################################
Erika Project - We share beautiful dreams together.
Brief: Erika is a AI develop by Go-lang.
Version: v1.00a
Author: alopex

Usage: erika [help] [start/stop] [status]

Options:
	help	- help information about erika application.
`)
	if err != nil {
		log.Println("Error print information:", err)
		os.Exit(1)
	}
	flag.PrintDefaults()
}
