package spider

import (
	"crypto/tls"
	"net/http"
)

var https *http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	https = &http.Client{
		Transport: tr,
	}
}
