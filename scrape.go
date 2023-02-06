package main

import (
	"fmt"

	"github.com/gocolly/colly"
	//"strconv"
)

type ForeclosureProperty struct {
	address       string
	price         string
	bed           string
	bath          string
	squareFootage string
}

// whatever is inside main runs
func main() {

	c := colly.NewCollector(colly.AllowedDomains("alachua.realforeclose.com"))

	/*for i := 1; i < 31; i++ {
	if i < 10{
		var date string = "0" + strconv.Itoa(i)
	} else {
		var date string = strconv.Itoa(i)
	}
	*/
	var date string = "06"

	scrapeURL := "https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&Zmethod=PREVIEW&AUCTIONDATE=02/" + date + "/2023"

	c.OnHTML("div.Head_C div[id=Area_C] div[tabindex=0] div.AUCTION_DETAILS tr", func(e *colly.HTMLElement) {

		fmt.Println(e.ChildText("td.AD_DTA"))
		//fmt.Println(e.ChildText("td.AD_DTA"))
	})

	c.Visit(scrapeURL)
}

//}
