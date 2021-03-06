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

func checkNetsUnpackParameters(t TNetsUnpack) (b bool, err error) {
	b = true
	// check src files
	if t.Src == "" {
		b = false
		log.Println("Source file can't be empty.")
		return b, err
	}
	b, err = PathExist(t.Src)
	if err != nil {
		log.Println("Error check path exist:", err)
		return b, err
	}
	if !b {
		log.Println("Source file path not exist.")
		return b, err
	}
	return b, err
}

func checkNetsPackProcessParameters(t TNetsPackProcessReq) (b bool, err error) {
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

func checkNetsUnpackVerboseParameters(t TNetsUnpackVerboseReq) (b bool, err error) {
	b = true
	// check src files
	if t.Src == "" {
		b = false
		log.Println("Source file can't be empty.")
		return b, err
	}
	b, err = PathExist(t.Src)
	if err != nil {
		log.Println("Error check path exist:", err)
		return b, err
	}
	if !b {
		log.Println("Source file path not exist.")
		return b, err
	}
	return b, err
}

func checkNetsUnpackProcessParameters(t TNetsUnpackProcessReq) (b bool, err error) {
	b = true
	// check src files
	if t.Src == "" {
		b = false
		log.Println("Source file can't be empty.")
		return b, err
	}
	b, err = PathExist(t.Src)
	if err != nil {
		log.Println("Error check path exist:", err)
		return b, err
	}
	if !b {
		log.Println("Source file path not exist.")
		return b, err
	}
	return b, err
}

func checkNetsUnpackToFileParameters(t TNetsUnpackToFile) (b bool, err error) {
	b = true
	// check src files
	if t.Src == "" {
		b = false
		log.Println("Source file can't be empty.")
		return b, err
	}
	b, err = PathExist(t.Src)
	if err != nil {
		log.Println("Error check path exist:", err)
		return b, err
	}
	if !b {
		log.Println("Source file path not exist.")
		return b, err
	}
	return b, err
}

func checkNetsUnpackToMemoryParameters(t TNetsUnpackToMemory) (b bool, err error) {
	b = true
	// check src files
	if t.Src == "" {
		b = false
		log.Println("Source file can't be empty.")
		return b, err
	}
	b, err = PathExist(t.Src)
	if err != nil {
		log.Println("Error check path exist:", err)
		return b, err
	}
	if !b {
		log.Println("Source file path not exist.")
		return b, err
	}
	return b, err
}

func checkNetsCompParameters(t TNetsComp) (b bool, err error) {
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
	case "TAR", "tar":
	case "TAR.GZ", "tar.gz":
	case "ZIP", "zip":
	default:
		b = false
		fmt.Printf("Algorithm %v not support.\n", t.Type)
	}
	return b, err
}

func checkNetsDecompParameters(t TNetsDecomp) (b bool, err error) {
	b = true
	// check src files
	if t.Src == "" {
		b = false
		log.Println("Source file can't be empty.")
		return b, err
	}
	b, err = PathExist(t.Src)
	if err != nil {
		log.Println("Error check path exist:", err)
		return b, err
	}
	if !b {
		log.Println("Source file path not exist.")
		return b, err
	}
	// check algorithm
	switch t.Type {
	case "TAR", "tar":
	case "TAR.GZ", "tar.gz":
	case "ZIP", "zip":
	default:
		b = false
		fmt.Printf("Algorithm %v not support.\n", t.Type)
	}
	return b, err
}

func checkNetsImagesQRCodeParameters(t TNetsImagesQRCodeToMemory) (b bool, err error) {
	b = true
	// check content
	if t.Content == "" {
		b = false
		log.Println("Content string can't be empty.")
		return b, err
	}
	// check size
	if t.Size <= 0 {
		b = false
		log.Println("Size should >= 0.")
		return b, err
	}
	return b, err
}

func checkNetsImagesQRCodeToFileParameters(t TNetsImagesQRCodeToFile) (b bool, err error) {
	b = true
	// check content
	if t.Content == "" {
		b = false
		log.Println("Content string can't be empty.")
		return b, err
	}
	// check size
	if t.Size <= 0 {
		b = false
		log.Println("Size should >= 0.")
		return b, err
	}
	return b, err
}

func checkNetsParsesIniValueParameters(t TNetsParsesIni) (b bool, err error) {
	b = true
	// check src files
	if t.Src == "" {
		b = false
		log.Println("Source file can't be empty.")
		return b, err
	}
	b, err = PathExist(t.Src)
	if err != nil {
		log.Println("Error check path exist:", err)
		return b, err
	}
	if !b {
		log.Println("Source file path not exist.")
		return b, err
	}
	// check parses mode
	if t.Mode != "get" && t.Mode != "set" {
		b = false
		log.Println("Mode should be one of 'get' or 'set'.")
	}
	// check parses type
	if t.Type != "string" && t.Type != "int" && t.Type != "float64" && t.Type != "bool" {
		b = false
		log.Println("Type should be one of 'string', 'int', 'float64' or 'bool'.")
	}
	// check parses value
	if t.Mode != "get" && t.Value == "" {
		b = false
		log.Println("Value should not be empty in set mode.")
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

func refactorNetsCompSource(src []string) (dest []string, err error) {
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
