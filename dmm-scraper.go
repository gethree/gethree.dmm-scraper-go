package main

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// Result is
type Result struct {
	PackageImageThumbURL string
	PackageImageURL      string
	ActressList          []*Actress
	DirectorList         []*Director
}

// Actress : Actress Info Struct
type Actress struct {
	ID   string
	Name string
}

// Director : Director Info Struct
type Director struct {
	ID   string
	Name string
}

func main() {
	//	doc, err := goquery.NewDocument("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=h_150upsm00217/")
	doc, err := goquery.NewDocument("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=hnds00028/")

	if err != nil {
		fmt.Print("url scarapping failed")
	}

	result := Result{}
	idMatcher := regexp.MustCompile(`id=([0-9]+)`)

	// タイトル画像
	doc.Find("#package-src-h_150upsm00217").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		result.PackageImageThumbURL = url
		fmt.Println(result.PackageImageThumbURL)
	})

	// 女優
	doc.Find("#performer > a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		content := s.Text()

		id := idMatcher.FindString(url)
		id = idMatcher.ReplaceAllString(id, "$1")

		actress := Actress{id, content}

		result.ActressList = append(result.ActressList, &actress)
	})

	fmt.Println(result.ActressList)

	// 監督
	doc.Find("a[href *= 'director']").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		content := s.Text()

		id := idMatcher.FindString(url)
		id = idMatcher.ReplaceAllString(id, "$1")
		director := Director{id, content}

		result.DirectorList = append(result.DirectorList, &director)
	})

	// シリーズ
	doc.Find("table.mg-b20 a[href *= 'article=series']").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		content := s.Text()
		fmt.Println(url)
		fmt.Println(content)
	})
}
