package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	. "satellite/global"
	"satellite/pack"
	. "satellite/utils"
	"sync/atomic"
	"time"

	"github.com/schollz/progressbar"
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

// ParseCmdPack function
// this function will be called in main.go and parse and execute one pack command
// pack command has three parameters
// input src file list, output dest file path and algorithm which used in unpack, return error info
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// algorithm now support 'aes', 'des', '3des', 'rsa', 'base64', you can send both up case and low case
// any failure or error will output print to screen and exit process
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

// handleCmdPack function
// this function mainly handle the main flow of command
// first check parameters legally, otherwise exit and return error
// then calculate the total work the pack should do, otherwise exit and return error
// start one go-routine to execute pack function and check result
// timeout will break and exit
func handleCmdPack(src []string, dest string, algorithm string) (err error) {
	ch := make(chan bool)
	// check parameters
	is := checkParameters(src, dest, algorithm)
	if !is {
		err = errors.New("parameters illegal")
		return err
	}
	src, err = refactorSource(src)
	if err != nil {
		fmt.Println("Error refactor source files:", err)
		return err
	}
	// calculate work
	var work int64
	err = pack.WorkCalculate(src, algorithm, &work)
	if err != nil {
		fmt.Println("Error calculate pack work")
		return err
	}
	if work <= 0 {
		err = errors.New("work can not equal or less zero")
		fmt.Println("Error work value")
		return err
	}
	fmt.Println("Pack Start:")
	// create process bar
	bar := progressbar.New64(work)
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
			/*for _, r := range "-\\|/" {
				fmt.Printf("\r%c, total:%d, done:%d", r, atomic.LoadInt64(&pack.NumAll), atomic.LoadInt64(&pack.NumDone))
				time.Sleep(100 * time.Millisecond)
			}*/
			done := atomic.LoadInt64(&pack.Done)
			atomic.StoreInt64(&pack.Done, 0)
			switch algorithm {
			case "AES", "aes":
				done *= AESBufferSize
			case "DES", "des":
				done *= DESBufferSize
			case "3DES", "3des":
				done *= DESBufferSize
			case "RSA", "rsa":
				done *= RSAPacketSize
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

// checkParameters function
// this function is mainly used to check parameters which from cli input and return check result
// check src files not empty and exist
// check algorithm is supported
// return check result(true/false)
func checkParameters(src []string, dest string, algorithm string) (is bool) {
	is = true
	// check src
	if len(src) == 0 {
		is = false
		fmt.Println("Source file list can't be empty.")
		return is
	}
	for i := 0; i < len(src); i++ {
		is, _ = PathExist(src[i])
		if !is {
			fmt.Printf("Source file %v path not exist.\n", i+1)
			return is
		}
	}
	// check algorithm
	switch algorithm {
	case "AES", "aes":
	case "DES", "des":
	case "3DES", "3des":
	case "RSA", "rsa":
	case "BASE64", "base64":
	default:
		is = false
		fmt.Printf("Algorithm %v not support.\n", algorithm)
	}
	return is
}

// refactorSource function
// this function is mainly used to refactor the source files path
func refactorSource(src []string) (dest []string, err error) {
	for i := 0; i < len(src); {
		is, err := IsDir(src[i])
		if err != nil {
			return dest, err
		}
		if is {
			err = filepath.Walk(src[i], func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return err
				}
				dest = append(dest, path)
				return err
			})
			if err != nil {
				return dest, err
			}
			src = append(src[:i], src[i+1:]...)
		} else {
			dest = append(dest, src[i])
			i++
		}
	}
	return dest, err
}

// execPack function
// this function is the process will be execute in new go-routine
func execPack(src []string, dest string, algorithm string, err *error, ch chan bool) {
	*err = pack.Pack(src, dest, algorithm)
	if *err != nil {
		ch <- false
		return
	}
	ch <- true
	return
}
