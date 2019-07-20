package pack

import (
	"encoding/base64"
	"sync"
)

/*func PackBase64One(srcfile string) (r string, err error) {
	// first, open the file
	file, err := os.Open(srcfile)
	if err != nil {
		log.Println("Error open file:", err)
		return r, err
	}
	defer file.Close()
	// second, read file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file:", err)
		return r, err
	}
}*/

func Base64EncrypteGo(str string, r *string, wg *sync.WaitGroup) {
	*r = Base64Encrypt(str)
	wg.Done()
}

func Base64Encrypt(str string) string {
	s := []byte(str)
	r := base64.StdEncoding.EncodeToString(s)
	return r
}
