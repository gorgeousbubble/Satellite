package pack

import (
	"github.com/pkg/errors"
	"log"
)

func Pack(srcfilelist []string, destfile string, algorithm string) (err error) {
	switch algorithm {
	case "AES", "aes":
		err = PackAES(srcfilelist, destfile)
	case "DES", "des":
		err = PackDES(srcfilelist, destfile)
	case "3DES", "3des":
		err = Pack3DES(srcfilelist, destfile)
	case "RSA", "rsa":
		err = PackRSA(srcfilelist, destfile)
	case "BASE64", "base64":
		err = PackBase64(srcfilelist, destfile)
	default:
		err = errors.New("Undefined unpack algorithm.")
		log.Println(err)
	}
	return err
}
