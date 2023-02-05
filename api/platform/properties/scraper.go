package properties

import (
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	//temp variable for location
	location string = "www.zillow.com/gainesville-fl"
)

func Connect() {
	d, err := gorm.Open("mysql", "go:Gators123@tcp(cen3031-project.mysql.database.azure.com)/listings?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func scrape() {
	c := colly.NewCollector(colly.AllowedDomains(location))

}
