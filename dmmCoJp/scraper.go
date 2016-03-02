package dmmCoJp

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeFromItemDetail :
func ScrapeFromItemDetail(itemDetailURL string) *DmmCoJpItem {

	doc, err := goquery.NewDocument(itemDetailURL)

	if err != nil {
		fmt.Print("url scarapping failed")
	}

	result := DmmCoJpItem{}
	cidMatcher := regexp.MustCompile(`cid=([^/]+)`)
	idMatcher := regexp.MustCompile(`id=([0-9]+)`)

	// 商品コード
	itemCode := cidMatcher.FindString(itemDetailURL)
	itemCode = cidMatcher.ReplaceAllString(itemCode, "$1")
	result.ItemCode = itemCode

	// タイトル
	doc.Find("#title").Each(func(_ int, s *goquery.Selection) {
		content := s.Text()
		result.Title = content
	})

	// タイトル画像
	doc.Find("#" + result.ItemCode).Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		result.PackageImageURL = url
	})

	// タイトル画像サムネイル
	doc.Find("#package-src-" + result.ItemCode).Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		result.PackageImageThumbURL = url
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

		id := idMatcher.FindString(url)
		id = idMatcher.ReplaceAllString(id, "$1")
		series := Series{id, content}

		result.SeriesList = append(result.SeriesList, &series)
	})

	// キーワード
	doc.Find("table.mg-b20 a[href *= 'article=keyword']").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		content := s.Text()

		id := idMatcher.FindString(url)
		id = idMatcher.ReplaceAllString(id, "$1")
		keyword := Keyword{id, content}

		result.KeywordList = append(result.KeywordList, &keyword)
	})

	// サンプル画像
	doc.Find("#sample-image-block img.mg-b6").Each(func(_ int, s *goquery.Selection) {
		thumbnailURL, _ := s.Attr("src")

		itemCodeMatcher := regexp.MustCompile(result.ItemCode + `-`)
		fullImageURL := itemCodeMatcher.ReplaceAllString(thumbnailURL, result.ItemCode+"jp-")

		sampleImage := SampleImage{thumbnailURL, fullImageURL}

		result.SampleImageList = append(result.SampleImageList, &sampleImage)
	})

	// 配信開始日/商品発売日
	doc.Find("table.mg-b20 td.nw").Each(func(_ int, s *goquery.Selection) {
		content := s.Text()
		if strings.Contains(content, "配信開始日") ||
			strings.Contains(content, "商品発売日") {
			fmt.Println(strings.TrimSpace(s.Next().Text()))
		}
	})

	return &result
}

// MarshalDmmCoJpItemToJSON :
func MarshalDmmCoJpItemToJSON(item *DmmCoJpItem) string {
	bytes, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return ""
	}
	return string(bytes)
}
