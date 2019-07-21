package unpack

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	. "satellite/utils"
	"strings"
	"sync"
)

func UnpackBase64One(data []byte, head TUnpackBase64One, destpath string) (err error) {
	// initial, fill the name
	var s []byte
	for _, v := range head.Name {
		if v == byte(0) {
			break
		}
		s = append(s, v)
	}
	destfile := destpath + string(s)
	// first, split the data slice
	ss, err := SplitByte(data, ConstBase64BufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return err
	}
	size := len(data) % ConstBase64BufferSize
	if size != 0 {
		last := len(data) / ConstBase64BufferSize
		ss[last] = append(ss[last][:0], ss[last][:size]...)
	}
	// second, we can call Base64Decrypt function
	wg := &sync.WaitGroup{}
	rr := make([]string, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go Base64DecryptGo(string(v), &rr[k], wg)
	}
	wg.Wait()
	dest := strings.Join(rr, "")
	// fourth, create the origin file
	err = ioutil.WriteFile(destfile, []byte(dest), 0644)
	if err != nil {
		log.Println("Error write to destfile:", err)
		return err
	}
	return err
}

func Base64DecryptGo(str string, r *string, wg *sync.WaitGroup) {
	*r = Base64Decrypt(str)
	wg.Done()
}

func Base64Decrypt(str string) string {
	s, _ := base64.StdEncoding.DecodeString(str)
	r := string(s)
	return r
}
