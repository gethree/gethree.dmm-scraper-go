package main

import (
	"fmt"

	"github.com/gethree/gethree.dmm-scraper-go/dmmCoJp"
)

func main() {
	result := dmmCoJp.ScrapeFromItemDetail("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=hnds00028/")
	fmt.Println(result.Title)
	fmt.Println(dmmCoJp.MarshalDmmCoJpItemToJSON(result))
	result = dmmCoJp.ScrapeFromItemDetail("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=h_150upsm00217/")
	fmt.Println(result.Title)
	fmt.Println(dmmCoJp.MarshalDmmCoJpItemToJSON(result))
	result = dmmCoJp.ScrapeFromItemDetail("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=snth00001/")
	fmt.Println(result.Title)
	fmt.Println(dmmCoJp.MarshalDmmCoJpItemToJSON(result))
}
