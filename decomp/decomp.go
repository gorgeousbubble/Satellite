package decomp

import (
	"github.com/pkg/errors"
)

func DeCompress(src string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = DeCompressTarGz(src, dest)
	case "zip":
		err = DeCompressZip(src, dest)
	default:
		err = errors.New("Undefined decompress algorithm.")
	}
	return err
}
