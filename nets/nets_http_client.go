package nets

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet(url string) (r []byte, err error) {
	// http Get request...
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error http get request:", err)
		return r, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("response status code not ok")
		log.Println("Error http response code:", err)
		return r, err
	}
	// start read response body...
	defer resp.Body.Close()
	r, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error read response body:", err)
		return r, err
	}
	return r, err
}

func HttpPost(url string, contentType string, body io.Reader) (r []byte, err error) {
	// http Post request...
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Println("Error http post request:", err)
		return r, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("response status code not ok")
		log.Println("Error http response code:", err)
		return r, err
	}
	// start read response body...
	defer resp.Body.Close()
	r, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error read response body:", err)
		return r, err
	}
	return r, err
}
