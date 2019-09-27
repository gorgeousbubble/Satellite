package nets

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	. "satellite/utils"
)

func checkNetsPackParameters(t TNetsPack) (b bool, err error) {
	b = true
	// check src files
	if t.Src == nil {
		b = false
		log.Println("Source file list can't be empty.")
		return b, err
	}
	for _, v := range t.Src {
		b, err = PathExist(v)
		if err != nil {
			log.Println("Error check path exist:", err)
			return b, err
		}
		if !b {
			log.Printf("Source file path not exist: '%v'\n", v)
			return b, err
		}
	}
	// check algorithm
	switch t.Type {
	case "AES", "aes":
	case "DES", "des":
	case "3DES", "3des":
	case "RSA", "rsa":
	case "BASE64", "base64":
	default:
		b = false
		fmt.Printf("Algorithm %v not support.\n", t.Type)
	}
	return b, err
}

func refactorNetsPackSource(src []string) (dest []string, err error) {
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
