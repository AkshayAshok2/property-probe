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
	fmt.Println("HI")
	var date string = "25"

	scrapeURL := "https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&Zmethod=PREVIEW&AUCTIONDATE=08/" + date + "/2022"

	c.OnHTML("div.Sub_Title", func(e *colly.HTMLElement) {

		fmt.Println(e.Text)

		//fmt.Println(e.ChildText(e.Text))
		//fmt.Println(e.ChildText("td.AD_DTA"))

	})

	c.OnHTML("div.Auct_Area", func(e *colly.HTMLElement) {
		selection := e.DOM
		fmt.Println(selection.Find("div.AUCTION_DETAILS > table > tbody > tr:nth-child(3) > td").Text())
	})

	c.OnXML("//tr //td", func(e *colly.XMLElement) {
		fmt.Println("ON XML: " + e.Text)
	})

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		fmt.Print("Childe: ")
		fmt.Println(e.ChildText)

	})

	c.Visit(scrapeURL)
}

//}
