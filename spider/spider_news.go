package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetNews(url string) (err error) {
	return err
}

func GetNewsShort(url string) (err error) {
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
	// find content news author
	doc.Find(".content p").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		val, exists := s.ChildrenFiltered("a").Attr("href")
		if exists {
			fmt.Println(url + val)
		}
	})
	return err
}
