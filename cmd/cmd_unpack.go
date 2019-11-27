package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/unpack"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/schollz/progressbar"
)

const line string = "------------------------------------------------------------------------------------"

var unpackCmd = flag.NewFlagSet(CmdUnpack, flag.ExitOnError)
var unpackSrc string
var unpackDest string
var unpackTarget string
var unpackVerbose bool
var unpackConfine bool

func init() {
	unpackCmd.StringVar(&unpackSrc, "i", "", "input files: packet file, such as \"file.dat\" or \"file.pak\"")
	unpackCmd.StringVar(&unpackDest, "o", "", "output files: one or more origin files. (should be path not file)")
	unpackCmd.StringVar(&unpackTarget, "t", "", "target file name: unpack choose one file. (should be file name)")
	unpackCmd.BoolVar(&unpackVerbose, "v", false, "verbose information list.")
	unpackCmd.BoolVar(&unpackConfine, "c", false, "unpack confine goroutine.")
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
	err = handleCmdUnpack(unpackSrc, unpackDest, unpackTarget, unpackVerbose, unpackConfine)
	if err != nil {
		fmt.Print("\n")
		fmt.Println("Unpack Failure:", err)
		os.Exit(1)
	}
	fmt.Print("\n")
	fmt.Println("Unpack Success.")
}

func handleCmdUnpack(src string, dest string, target string, verbose bool, confine bool) (err error) {
	var algorithm string
	ch := make(chan bool)
	// whether look up verbose information
	if verbose {
		var files []string
		var sizes []int
		err = unpack.ExtractInfo(src, &files, &sizes, &algorithm)
		if err != nil {
			fmt.Println("Error Extract Unpack Information")
			return err
		}
		fmt.Println(line)
		fmt.Printf("%-32s\t%-16s\t%-16s\n", "file", "size", "type")
		fmt.Println(line)
		for i := 0; i < len(files); i++ {
			name := files[i]
			size := sizes[i] / 1024
			fmt.Printf("%-32s\t%-16s\t%-16s\n", name, strconv.Itoa(size)+"kb", algorithm)
		}
		fmt.Println(line)
		return nil
	}
	// calculate work
	var work int64
	err = unpack.WorkCalculate(src, &algorithm, &work)
	if err != nil || work <= 0 {
		fmt.Println("Error Calculate Unpack Work")
		return err
	}
	fmt.Println("Unpack Start:")
	// create process bar
	bar := progressbar.New64(work)
	// execute unpack function
	go execUnpack(src, dest, target, confine, &err, ch)
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

func execUnpack(src string, dest string, target string, confine bool, err *error, ch chan bool) {
	if target != "" {
		*err = unpack.UnpackToFile(src, target, dest)
	} else {
		if confine {
			*err = unpack.UnpackConfine(src, dest)
		} else {
			*err = unpack.Unpack(src, dest)
		}
	}
	if *err != nil {
		ch <- false
		return
	}
	ch <- true
	return
}
