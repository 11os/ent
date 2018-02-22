package utils

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// Run ...
func Run(url string) (obj []string) {
	reader, _ := get(url)
	node, err := html.Parse(reader)
	doc := goquery.NewDocumentFromNode(node)
	if err != nil {
		log.Fatal(err)
	}
	title, _ := doc.Find("title").Html()
	print(title)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		hrefGbk, _ := decodeToGBK(href)
		if strings.Contains(hrefGbk, "ftp") {
			// print(href)
			obj = append(obj, hrefGbk)
			// obj = obj.append(hrefGbk)
		}
	})
	return
}
