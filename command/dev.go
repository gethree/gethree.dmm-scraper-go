package command

import (
	"fmt"
	"strings"

	"github.com/gethree/gethree.dmm-scraper-go/dmmCoJp"
)

type DevCommand struct {
	Meta
}

func (c *DevCommand) Run(args []string) int {
	result := dmmCoJp.ScrapeFromItemDetail("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=hnds00028/")
	fmt.Println(result.Title)
	fmt.Println(dmmCoJp.MarshalDmmCoJpItemToJSON(result))
	result = dmmCoJp.ScrapeFromItemDetail("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=h_150upsm00217/")
	fmt.Println(result.Title)
	fmt.Println(dmmCoJp.MarshalDmmCoJpItemToJSON(result))
	result = dmmCoJp.ScrapeFromItemDetail("http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=snth00001/")
	fmt.Println(result.Title)
	fmt.Println(dmmCoJp.MarshalDmmCoJpItemToJSON(result))

	return 0
}

func (c *DevCommand) Synopsis() string {
	return "dummy command for dev"
}

func (c *DevCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
