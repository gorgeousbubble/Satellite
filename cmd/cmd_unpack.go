package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/unpack"
	"sync/atomic"
	"time"

	"github.com/schollz/progressbar"
)

const line string = "--------------------------------------------------"

var unpackCmd = flag.NewFlagSet(CmdUnpack, flag.ExitOnError)
var unpackSrc string
var unpackDest string
var unpackVerbose bool

func init() {
	unpackCmd.StringVar(&unpackSrc, "i", "", "input files: packet file, such as \"file.dat\" or \"file.pak\"")
	unpackCmd.StringVar(&unpackDest, "o", "", "output files: one or more origin files. (should be path not file)")
	unpackCmd.BoolVar(&unpackVerbose, "v", false, "verbose information list.")
}

func ParseCmdUnpack() {
	// check args number
	if len(os.Args) == 2 {
		unpackCmd.Usage()
		os.Exit(1)
	}
	// parse command unpack
	err := unpackCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Unpack Command.")
		os.Exit(1)
	}
	// handle command parameters
	err = handleCmdUnpack(unpackSrc, unpackDest, unpackVerbose)
	if err != nil {
		fmt.Print("\n")
		fmt.Println("Unpack Failure:", err)
		os.Exit(1)
	}
	fmt.Print("\n")
	fmt.Println("Unpack Success.")
}

func handleCmdUnpack(src string, dest string, verbose bool) (err error) {
	// whether look up verbose information
	if verbose {
		var info []string
		err = unpack.ExtractInfo(src, &info)
		if err != nil {
			fmt.Println("Error Extract Unpack Information")
			return err
		}
		fmt.Println(line)
		fmt.Println("Files")
		fmt.Println(line)
		for _, v := range info {
			fmt.Println(v)
		}
		fmt.Println(line)
		return nil
	}
	ch := make(chan bool)
	// calculate work
	var work int64
	var algorithm string
	err = unpack.WorkCalculate(src, &algorithm, &work)
	if err != nil || work <= 0 {
		fmt.Println("Error Calculate Unpack Work")
		return err
	}
	fmt.Println("Unpack Start:")
	// create process bar
	bar := progressbar.New64(work)
	// execute unpack function
	go execUnpack(src, dest, &err, ch)
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Unpack failure:", err)
				return err
			}
			log.Println("Unpack success.")
			return err
		default:
			done := atomic.LoadInt64(&unpack.Done)
			atomic.StoreInt64(&unpack.Done, 0)
			switch algorithm {
			case "AES", "aes":
				done *= AESBufferSize
			case "DES", "des":
				done *= DESBufferSize
			case "3DES", "3des":
				done *= DESBufferSize
			case "RSA", "rsa":
				done *= RSAUnpackSize
			case "BASE64", "base64":
				done *= Base64BufferSize
			default:
				s := fmt.Sprint("Undefined pack algorithm.")
				err = errors.New(s)
				return err
			}
			err = bar.Add64(done)
			if err != nil {
				fmt.Println("Error add count:", err)
				return err
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func execUnpack(src string, dest string, err *error, ch chan bool) {
	*err = unpack.Unpack(src, dest)
	if *err != nil {
		ch <- false
		return
	}
	ch <- true
	return
}
