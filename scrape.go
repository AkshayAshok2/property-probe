package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Property struct {
	AuctionType     string
	JudgmentAmount  float64
	PropertyAddress string
	AssessedValue   float64
}

func main() {
	// Open a connection to the MySQL database using GORM
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Auto-migrate the Property struct to create the 'properties' table if it doesn't exist
	db.AutoMigrate(&Property{})

	// Define a new Property object with sample data
	prop := Property{
		AuctionType:     "foreclosure",
		JudgmentAmount:  100000.0,
		PropertyAddress: "123 Main St",
		AssessedValue:   80000.0,
	}

	// Create a new record in the 'properties' table using GORM's Create() function
	result := db.Create(&prop)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Successfully inserted property data into MySQL database")
}

/*
func main() {
	// Create a new Colly collector
	c := colly.NewCollector()

	// Set the URL of the webpage to scrape
	url := "https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&zmethod=PREVIEW&AuctionDate=09/20/2022"

	// Set the headers for the request
	headers := map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "en-US,en;q=0.9",
		"Referer":         "https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&zmethod=PREVIEW&AuctionDate=09/20/2022",
		"Connection":      "keep-alive",
	}

	// Before making the request, set the custom headers
	c.OnRequest(func(r *colly.Request) {
		for k, v := range headers {
			r.Headers.Set(k, v)
		}
	})

	// Find the target element using a CSS selector
	c.OnHTML("#AITEM_1336448 > div.AUCTION_DETAILS > table > tbody > tr:nth-child(2) > td", func(e *colly.HTMLElement) {
		// Extract the text content of the element
		elementText := e.Text

		// Print the extracted text
		fmt.Println(elementText)
	})

	// Visit the URL
	c.Visit(url)
}

/*
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ForeclosureProperty struct {
	address       string `json: "Address"`
	price         string `json: "Price"`
	bed           string `json: "Bed"`
	bath          string `json: "Bath"`
	squareFootage string `json: "Square Footage"`
}

func connectAndLoadData(w http.ResponseWriter, r *http.Request) {

	property := ForeclosureProperty{
		address:       "123",
		price:         "1000",
		bed:           "3",
		bath:          "4",
		squareFootage: "2000",
	}
	var dsn = "go:Gators123@tcp(cen3031-project.mysql.database.azure.com)/listings?charset=utf8&parseTime=True&Ioc=Local"
	var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Create(property)
	if err := db.Create(property).Error; err != nil {
		log.Fatalln((err))
	}

	json.NewEncoder(w).Encode(property)
	fmt.Println("Fields Added", property)

}

// whatever is inside main runs
func main() {
	/*http.HandleFunc("/CreateStuff",connectAndLoadData)
	log.Fatal(http.ListenAndServe(":8080", nil))
	c := colly.NewCollector(colly.AllowedDomains("alachua.realforeclose.com"))
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", r.URL)
	})

	//Initializes the days and strings to loop through the urls and gather foreclosure data for each day
	//var month string = "02"
	//var day string = "09"

	/*for i := 1; i < 13; i++ {
		if i < 10{
			var month string = "0" + strconv.Itoa(i)
		} else {
			var month string = strconv.Itoa(i)
		}

		for j := 1; j < 28; i++ {
			if j < 10{
				var day string = "0" + strconv.Itoa(i)
			} else {
				var day string = strconv.Itoa(i)
			}
		}
	}

	var month string = "09"
	var day string = "20"

	scrapeURL := "https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&Zmethod=PREVIEW&AUCTIONDATE=" + month + "/" + day + "/2022"

	//c.OnHTML("div.Auct_Area", func(e *colly.HTMLElement) {
	//	selection := e.DOM
	//	fmt.Println(selection.Find("div.AUCTION_DETAILS > table > tbody > tr:nth-child(3) > td").Text())
	//})

	c.OnHTML("div.AUCTION_DETAILS > table > tbody > tr:nth-child(1) > td" /*"body.main_text.div.headerMenu" /*body.main_text.table.MAIN_TBL.tbody" /*.tr.td.MAIN_TBL_CONTENT", func(e *colly.HTMLElement) {

		var contentHeader string = e.ChildText("th")

		fmt.Println(contentHeader)
	})
	//})
	// Visit link found on page
	// Only those links are visited which are in AllowedDomains
	c.Visit(scrapeURL)

	var scrapeAddress string
	var scrapeBath string
	var scrapeBed string
	var scrapePrice string
	var scrapeSquareFootage string

	var property ForeclosureProperty
	property.address = scrapeAddress
	property.bath = scrapeBath
	property.bed = scrapeBed
	property.price = scrapePrice
	property.squareFootage = scrapeSquareFootage
}

/*
   	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "go", "Gators123", host, port, database)
   	// Open the connection
   	db, err := sql.Open("mysql", dsn)
   	if err != nil {
       		log.Fatalf("impossible to create the connection: %s", err)
   	}
   	defer db.Close()
*/

//}
