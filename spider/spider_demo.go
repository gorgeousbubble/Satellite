package spider

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
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

func GetMovie(url string) (err error) {
	// combine http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// append request header
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	// send request and get response
	resp, err := https.Do(req)
	if err != nil {
		log.Println("Error send GET request:", err)
		return err
	}
	// check response status code
	if resp.StatusCode != http.StatusOK {
		log.Println("Error response status code:", err)
		return err
	}
	// create document from response body
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error create document from response body:", err)
		return err
	}
	// find movie info in response body
	doc.Find("#content h1").Each(func(i int, s *goquery.Selection) {
		// find name
		fmt.Println("name:" + s.ChildrenFiltered(`[property="v:itemreviewed"]`).Text())
		// find year
		fmt.Println("year:" + s.ChildrenFiltered(`.year`).Text())
	})
	// find director info
	director := ""
	doc.Find("#info span:nth-child(1) span.attrs").Each(func(i int, s *goquery.Selection) {
		director += s.Text()
	})
	fmt.Println("director:" + director)
	// find screenwriter info
	screenwriter := ""
	doc.Find("#info span:nth-child(3) span.attrs").Each(func(i int, s *goquery.Selection) {
		screenwriter += s.Text()
	})
	fmt.Println("screenwriter:" + screenwriter)
	// find character info
	character := ""
	doc.Find("#info span.actor span.attrs").Each(func(i int, s *goquery.Selection) {
		character += s.Text()
	})
	fmt.Println("character:" + character)
	// find category info
	category := ""
	doc.Find("#info > span:nth-child(8)").Each(func(i int, s *goquery.Selection) {
		category += s.Text()
	})
	fmt.Println("category:" + category)
	return err
}

func GetTopList(url string) (l []string, err error) {
	// combine http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// append request header
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	// send request and get response
	resp, err := https.Do(req)
	if err != nil {
		log.Println("Error send GET request:", err)
		return l, err
	}
	// check response status code
	if resp.StatusCode != http.StatusOK {
		log.Println("Error response status code:", err)
		return l, err
	}
	// create document from response body
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error create document from response body:", err)
		return l, err
	}
	// find movies link info
	doc.Find("#content div div.article ol li div div.info div.hd a").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%v", s)
		herf, _ := s.Attr("href")
		l = append(l, herf)
	})
	return l, err
}
