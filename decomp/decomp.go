package decomp

import (
	"github.com/pkg/errors"
	"log"
)

func DeCompress(src string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = DeCompressTarGz(src, dest)
	case "zip":
		err = DeCompressZip(src, dest)
	default:
		err = errors.New("Undefined decompress algorithm.")
		log.Println(err)
	}
	return err
}
