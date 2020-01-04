package nets

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var httpsClient *http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	httpsClient = &http.Client{
		Transport:tr,
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
