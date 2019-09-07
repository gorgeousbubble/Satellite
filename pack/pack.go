package pack

import (
	"errors"
	"fmt"
)

func Pack(src []string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "AES", "aes":
		err = PackAES(src, dest)
	case "DES", "des":
		err = PackDES(src, dest)
	case "3DES", "3des":
		err = Pack3DES(src, dest)
	case "RSA", "rsa":
		err = PackRSA(src, dest)
	case "BASE64", "base64":
		err = PackBase64(src, dest)
	default:
		s := fmt.Sprint("Undefined pack algorithm.")
		err = errors.New(s)
	}
	return err
}

func WorkCalculate(src []string, algorithm string, work *int64) (err error) {
	switch algorithm {
	case "AES", "aes":
		*work, err = PackAESWorkCalculate(src)
	case "DES", "des":
		*work, err = PackDESWorkCalculate(src)
	case "3DES", "3des":
		*work, err = PackDESWorkCalculate(src)
	case "RSA", "rsa":
		*work, err = PackRSAWorkCalculate(src)
	case "BASE64", "base64":
		*work, err = PackBase64WorkCalculate(src)
	default:
		s := fmt.Sprint("Undefined pack algorithm.")
		err = errors.New(s)
	}
	return err
}
