package utils

import (
	"fmt"
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

// Get58list ...
func Get58list(url string) (obj [][]map[string]string) {
	reader, _ := get(url)
	node, err := html.Parse(reader)
	doc := goquery.NewDocumentFromNode(node)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		pos, _ := s.Attr("_pos")
		// hrefGbk, _ := decodeToGBK(href)
		if pos != "" && count < 3 {
			count++
			fmt.Println("count = ", count)
			url, _ := s.Find(".title").Find("a").Attr("href")
			// picture, _ := s.Find(".pic").Find("a").Attr("href")
			// title := s.Find(".title").Find("a").Text()
			// sum, _ := s.Find(".sum").Html()
			// unit, _ := s.Find(".unit").Html()
			// price, _ := s.Find(".price").Html()
			// temp := map[string]string{
			// 	"url": url,
			// 	"position": pos,
			// 	"picture": picture,
			// 	"title": title,
			// 	"sum": sum,
			// 	"unit": unit,
			// 	"price": price,
			// "detail": Get58Detail(url),
			// }
			// obj = append(obj, temp)
			obj = append(obj, Get58Detail(url))
			// obj = Get58Detail(url)
		}

		// if strings.Contains(hrefGbk, "") {
		// print(href)

		// obj = obj.append(hrefGbk)
		// }
	})
	return
}

// Get58Detail ...
func Get58Detail(url string) (obj []map[string]string) {
	fmt.Println("start: ", url)
	reader, _ := get(url)
	node, err := html.Parse(reader)
	doc := goquery.NewDocumentFromNode(node)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".main-wrap").Each(func(i int, s *goquery.Selection) {
		// hrefGbk, _ := decodeToGBK(href)
		picture, _ := s.Find(".pic").Find("a").Attr("href")
		xtu1, _ := s.Find("#xtu_1").Attr("data-value")
		xtu2, _ := s.Find("#xtu_2").Attr("data-value")
		xtu3, _ := s.Find("#xtu_3").Attr("data-value")
		xtu4, _ := s.Find("#xtu_4").Attr("data-value")
		title := s.Find(".house-title").Find("h1").Text()
		sum := s.Find(".sum").Text()
		unit := s.Find(".unit").Text()
		price := s.Find(".price").Text()
		room_main := s.Find(".room").Find(".main").Text()
		room_sub := s.Find(".room").Find(".sub").Text()
		area_main := s.Find(".area").Find(".main").Text()
		area_sub := s.Find(".area").Find(".sub").Text()
		toward_main := s.Find(".toward").Find(".main").Text()
		toward_sub := s.Find(".toward").Find(".sub").Text()
		houseBasicInfo := s.Find("#houseBasicInfo").Find("#basicInfo").First().Text()
		temp := map[string]string{
			"url":            url,
			"picture":        picture,
			"title":          title,
			"sum":            sum,
			"unit":           unit,
			"price":          price,
			"xtu1":           xtu1,
			"xtu2":           xtu2,
			"xtu3":           xtu3,
			"xtu4":           xtu4,
			"room_main":      room_main,
			"room_sub":       room_sub,
			"area_main":      area_main,
			"area_sub":       area_sub,
			"toward_main":    toward_main,
			"toward_sub":     toward_sub,
			"houseBasicInfo": houseBasicInfo,
		}
		obj = append(obj, temp)
	})
	return
}
