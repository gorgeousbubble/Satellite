package shell

import (
	"errors"
	"fmt"
)

func Shell(src string, dest string, algorithm string) (r string, err error) {
	switch algorithm {
	case "upx":
		r, err = shellUpx(src, dest)
	default:
		s := fmt.Sprint("Undefined shell algorithm.")
		err = errors.New(s)
	}
	return r, err
}
