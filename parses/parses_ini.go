package parses

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
)

func getValue(in []byte, section string, key string) (r []byte) {
	// split data by '\n'
	var s [][]byte
	s1 := bytes.Split(in, []byte("\n"))
	for _, v := range s1 {
		// delete comments
		if bytes.Contains(v, []byte(";")) {
			index := bytes.Index(v, []byte(";"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete all space
		v = bytes.ReplaceAll(v, []byte(" "), []byte(""))
		// delete all return
		v = bytes.ReplaceAll(v, []byte("\r"), []byte(""))
		// delete all blank lines, packet to s slice
		if !bytes.Equal(v, []byte("")) {
			s = append(s, v)
		}
	}
	// seek for section
	var name []byte
	for _, v := range s {
		// section start with '[' and end with ']'
		if v[0] == '[' && v[len(v)-1] == ']' {
			name = v[1 : len(v)-1]
		} else if string(name) == section {
			// seek for key
			pairs := bytes.Split(v, []byte("="))
			if len(pairs) == 2 {
				// key value before '='
				k := bytes.TrimSpace(pairs[0])
				if string(k) == key {
					r = bytes.TrimSpace(pairs[1])
					return r
				}
			}
		}
	}
	r = []byte("")
	return r
}

func setValue(in []byte, section string, key string, value string) (r []byte, err error) {
	// split data by '\n'
	var s [][]byte
	var s2 [][]byte
	s1 := bytes.Split(in, []byte("\n"))
	for _, v := range s1 {
		// delete comments
		if bytes.Contains(v, []byte(";")) {
			index := bytes.Index(v, []byte(";"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete all space
		v = bytes.ReplaceAll(v, []byte(" "), []byte(""))
		// delete all return
		v = bytes.ReplaceAll(v, []byte("\r"), []byte(""))
		// delete all blank lines, packet to s slice
		if !bytes.Equal(v, []byte("")) {
			s2 = append(s2, v)
		}
	}
	// seek for section
	var flag = false
	var name []byte
	for _, v := range s2 {
		// section start with '[' and end with ']'
		if v[0] == '[' && v[len(v)-1] == ']' {
			name = v[1 : len(v)-1]
		} else if string(name) == section {
			// seek for key
			pairs := bytes.Split(v, []byte("="))
			if len(pairs) == 2 {
				// key value before '='
				k := bytes.TrimSpace(pairs[0])
				if string(k) == key {
					v = []byte(key + "=" + value)
					flag = true
				}
			}
		}
		s = append(s, v)
	}
	// check finished set value
	if !flag {
		err = errors.New("can not find key")
		return r, err
	}
	// splice slice
	r = bytes.Join(s, []byte("\n"))
	return r, err
}

func getValueFrom(src string, section string, key string) (r string, err error) {
	// open ini file...
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open ini:", err)
		return r, err
	}
	defer file.Close()
	// read ini file...
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read ini:", err)
		return r, err
	}
	// get key-value from stream
	out := getValue(data, section, key)
	if bytes.Equal(out, []byte("")) {
		err = errors.New("section or key may not exist")
		log.Println("Error get value:", err)
		return r, err
	}
	r = string(out)
	return r, err
}

func setValueTo(src string, section string, key string, value string) (err error) {
	// open ini file...
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open ini:", err)
		return err
	}
	// read ini file...
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read ini:", err)
		return err
	}
	// close ini file...
	err = file.Close()
	if err != nil {
		log.Println("Error close ini:", err)
		return err
	}
	// set key-value to stream
	out, err := setValue(data, section, key, value)
	if err != nil {
		log.Println("Error set value to ini:", err)
		return err
	}
	// write ini file...
	err = ioutil.WriteFile(src, out, 0644)
	if err != nil {
		log.Println("Error write ini:", err)
		return err
	}
	return err
}
