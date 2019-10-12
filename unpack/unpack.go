package unpack

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

func Unpack(src string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAES(src, dest)
	case "DES", "des":
		err = UnpackDES(src, dest)
	case "3DES", "3des":
		err = Unpack3DES(src, dest)
	case "RSA", "rsa":
		err = UnpackRSA(src, dest)
	case "BASE64", "base64":
		err = UnpackBase64(src, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

func UnpackToFile(src string, target string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAESToFile(src, target, dest)
	case "DES", "des":
		err = UnpackDESToFile(src, target, dest)
	case "3DES", "3des":
		err = Unpack3DESToFile(src, target, dest)
	case "RSA", "rsa":
		err = UnpackRSAToFile(src, target, dest)
	case "BASE64", "base64":
		err = UnpackBase64ToFile(src, target, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

func UnpackToMemory(src string, target string, dest *[]byte) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAESToMemory(src, target, dest)
	case "DES", "des":
		err = UnpackDESToMemory(src, target, dest)
	case "3DES", "3des":
		err = Unpack3DESToMemory(src, target, dest)
	case "RSA", "rsa":
		err = UnpackRSAToMemory(src, target, dest)
	case "BASE64", "base64":
		err = UnpackBase64ToMemory(src, target, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

func ExtractInfo(src string, dest *[]string, sz *[]int, algorithm *string) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAESExtractInfo(src, dest, sz)
		*algorithm = "aes"
	case "DES", "des":
		err = UnpackDESExtractInfo(src, dest, sz)
		*algorithm = "des"
	case "3DES", "3des":
		err = Unpack3DESExtractInfo(src, dest, sz)
		*algorithm = "3des"
	case "RSA", "rsa":
		err = UnpackRSAExtractInfo(src, dest, sz)
		*algorithm = "rsa"
	case "BASE64", "base64":
		err = UnpackBase64ExtractInfo(src, dest, sz)
		*algorithm = "base64"
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

func WorkCalculate(src string, algorithm *string, work *int64) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		*work, err = UnpackAESWorkCalculate(src)
		*algorithm = "AES"
	case "DES", "des":
		*work, err = UnpackDESWorkCalculate(src)
		*algorithm = "DES"
	case "3DES", "3des":
		*work, err = Unpack3DESWorkCalculate(src)
		*algorithm = "3DES"
	case "RSA", "rsa":
		*work, err = UnpackRSAWorkCalculate(src)
		*algorithm = "RSA"
	case "BASE64", "base64":
		*work, err = UnpackBase64WorkCalculate(src)
		*algorithm = "BASE64"
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}
