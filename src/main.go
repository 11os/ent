package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mozillazg/request"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func get(url string) (reader io.Reader, err error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = map[string]string{
		"Accept-Encoding": "",
	}
	resp, err := req.Get(url)
	if err != nil {
		return
	}
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
	} else {
		reader = resp.Body
	}

	return
}

func run(reader io.Reader) {
	node, err := html.Parse(reader)
	doc := goquery.NewDocumentFromNode(node)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		hrefGbk, _ := decodeToGBK(href)
		if strings.Contains(hrefGbk, "ftp") {
			fmt.Println(hrefGbk)
		}
	})
}

func decodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

func main() {
	reader, _ := get("https://www.dygod.net/html/tv/oumeitv/109955.html")
	run(reader)
}
