package comp

import (
	"github.com/pkg/errors"
	"log"
)

func Compress(src []string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = CompressTarGz(src, dest)
	case "zip":
		err = CompressZip(src, dest)
	default:
		err = errors.New("Undefined compress algorithm.")
		log.Println(err)
	}
	return err
}
