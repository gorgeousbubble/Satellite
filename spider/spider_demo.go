package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetMovie(url string) (err error) {
	// send request and get response
	resp, err := http.Get(url)
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
