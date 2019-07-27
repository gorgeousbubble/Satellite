package pack

import (
	"github.com/pkg/errors"
	"log"
)

func Pack(srcfilelist []string, destfile string, algorithm string) (err error) {
	switch algorithm {
	case "AES":
		err = PackAES(srcfilelist, destfile)
	case "DES":
		err = PackDES(srcfilelist, destfile)
	case "3DES":
		err = Pack3DES(srcfilelist, destfile)
	case "RSA":
		err = PackRSA(srcfilelist, destfile)
	case "BASE64":
		err = PackBase64(srcfilelist, destfile)
	default:
		err = errors.New("Undefined unpack algorithm.")
		log.Println(err)
	}
	return err
}
