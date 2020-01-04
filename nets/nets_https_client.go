package nets

import (
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var httpsClient *http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	httpsClient = &http.Client{
		Transport: tr,
	}
}

func HttpsGet(url string) (r []byte, err error) {
	// https Get request...
	resp, err := httpsClient.Get(url)
	if err != nil {
		log.Println("Error https get request:", err)
		return r, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("response status code not ok")
		log.Println("Error https response code:", err)
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

func HttpsPost(url string, contentType string, body io.Reader) (r []byte, err error) {
	// https Post request...
	resp, err := httpsClient.Post(url, contentType, body)
	if err != nil {
		log.Println("Error https post request:", err)
		return r, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("response status code not ok")
		log.Println("Error https response code:", err)
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

func HttpsPostForm(url string, data url.Values) (r []byte, err error) {
	// https PostFrom request...
	resp, err := httpsClient.PostForm(url, data)
	if err != nil {
		log.Println("Error https post form request:", err)
		return r, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("response status code not ok")
		log.Println("Error https response code:", err)
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

func HttpsHead(url string) (code int, err error) {
	// https Head request...
	resp, err := httpsClient.Head(url)
	if err != nil {
		log.Println("Error https head request:", err)
		return code, err
	}
	// return https status code
	code = resp.StatusCode
	return code, err
}
