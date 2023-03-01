package main

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Add(db *gorm.DB, property Property) {
	result := db.Create(property)
	if result.Error != nil {
		panic(result.Error)
	}

}
func getSize(db *gorm.DB) int {
	var properties []Property
	resultCheckFirst := db.Find(&properties)
	if resultCheckFirst.Error != nil {
		panic(resultCheckFirst.Error)
	}
	return len(properties)
}

func TestAdd(t *testing.T) {
	//Size before add function
	var sizeInitial = getSize(Connect())
	//Creates example property to add
	property := Property{
		auction_type:     "auction_type_1",
		judgement_amount: 1000.0,
		address:          "123 Main St",
		assessed_value:   5000.0,
	}
	// Insert the property into the database
	Add(Connect(), property)
	var sizeAfter = getSize(Connect())
	if sizeAfter <= sizeInitial {
		t.Errorf("No property was added")
	} else {
		fmt.Println("Property Successfully Added")
	}

}

/*func printPropertiesTableDetails(db *gorm.DB) {

	var properties []Property
	resultCheck := db.Find(&properties)
	if resultCheck.Error != nil {
		panic(resultCheck.Error)
	}

	for _, property := range properties {
		fmt.Println("ID: %d, AuctionType: %s, JudgementAmount: %f, PropertyAddress: %s, AssessedValue: %f\n",
			property.auction_type, property.judgement_amount, property.address, property.assessed_value)
	}
}*/

type Property struct {
	auction_type     string
	judgement_amount float64
	address          string
	assessed_value   float64
}

func main() {
	TestAdd(&testing.T{})

	/*property := Property{
			auction_type:     "auction_type_1",
			judgement_amount: 1000.0,
			address:          "123 Main St",
			assessed_value:   5000.0,
		}
		Add(property)
		printPropertiesTableDetails()
	}*/
}
