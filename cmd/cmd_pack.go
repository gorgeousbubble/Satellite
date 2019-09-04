package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/pack"
	"time"
)

var packCmd = flag.NewFlagSet(CmdPacket, flag.ExitOnError)
var packSrc []string
var packDest string
var packType string

func init() {
	packCmd.Var(NewStrSlice([]string{}, &packSrc), "i", "input files: file list support with different type, such as \"file_1.txt,file_2.mov,file_3.png...\"")
	packCmd.StringVar(&packDest, "o", "", "output files: one file which user can customize it type, such as \"file.dat\" or \"file.pak\"")
	packCmd.StringVar(&packType, "t", "AES", "pack type: one type of enum [AES,DES,3DES,RSA,BASE64]")
}

func ParseCmdPack() {
	// check args number
	if len(os.Args) == 2 {
		packCmd.Usage()
		os.Exit(1)
	}
	// parse command pack
	err := packCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Pack Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdPack(packSrc, packDest, packType)
	if err != nil {
		fmt.Print("\n")
		fmt.Println("Pack failure:", err)
		os.Exit(1)
	}
	fmt.Print("\n")
	fmt.Println("Pack success.")
}

func handleCmdPack(src []string, dest string, algorithm string) (err error) {
	ch := make(chan bool)
	// execute pack function
	go execPack(src, dest, algorithm, &err, ch)
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Pack failure:", err)
				return err
			}
			log.Println("Pack success.")
			return err
		default:
			for _, r := range "-\\|/" {
				fmt.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func execPack(src []string, dest string, algorithm string, err *error, ch chan bool) {
	*err = pack.Pack(src, dest, algorithm)
	if *err != nil {
		ch <- false
		return
	}
	ch <- true
	return
}
