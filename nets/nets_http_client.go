package nets

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func HttpPostForm(url string, data url.Values) (r []byte, err error) {
	// http PostFrom request...
	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Println("Error http post form request:", err)
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

func HttpHead(url string) (code int, err error) {
	// http Head request...
	resp, err := http.Head(url)
	if err != nil {
		log.Println("Error http head request:", err)
		return code, err
	}
	// return http status code
	code = resp.StatusCode
	return code, err
}
