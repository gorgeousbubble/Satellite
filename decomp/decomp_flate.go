package decomp

import (
	"bytes"
	"compress/flate"
	"io/ioutil"
)

func DeCompressFlate(src []byte) (dest []byte, err error) {
	// create new buffer from input
	in := bytes.NewBuffer(src)
	// create new flate reader
	fr := flate.NewReader(in)
	defer fr.Close()
	// copy to dest slice
	dest, _ = ioutil.ReadAll(fr)
	return dest, err
}
