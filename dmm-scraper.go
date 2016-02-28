package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=h_150upsm00217/")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("#package-src-h_150upsm00217").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		fmt.Println(url)
	})
}
