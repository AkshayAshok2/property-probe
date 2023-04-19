package properties

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
	"gorm.io/gorm"
)

type Property struct {
	Date            string  `json:"date"`
	AuctionType     string  `json:"auction_type"`
	JudgementAmount float64 `json:"judgement_amount"`
	Address         string  `json:"address"`
	AssessedValue   float64 `json:"assessedvalue"`
	LatLon          string  `json:"latlon"`
	Description     string  `json:"description"`
	ZipCode         string  `json:"zip_code"`
}

func CreateProperty(db *gorm.DB, property *Property) (err error) {
	err = db.Create(property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllProperties(db *gorm.DB, property *[]Property) (err error) {
	err = db.Find(property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProperty(db *gorm.DB, property *Property, address string) (err error) {
	err = db.Where("address = ?", address).First(property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetZipCodeProperties(db *gorm.DB, property *[]Property, zipcode string) (err error) {
	err = db.Where("zip_code = ?", zipcode).Find(property).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProperty(db *gorm.DB, property *Property) (err error) {
	db.Save(property)
	return nil
}

func DeleteProperty(db *gorm.DB, property *Property, address string) (err error) {
	db.Where("address = ?", address).Delete(property)
	return nil
}

func GetDescription(address string) (description string) {
	ctx := context.Background()
	if len(address) == 0 {
		return ""
	}

	results, err := googlesearch.Search(ctx, address+" zillow")
	if err != nil {
		return "No information on property found!"
	}

	if len(results) > 0 {
		for _, value := range results {
			if strings.HasPrefix(value.URL, "https://www.zillow.com") {
				description = value.Description
				break
			}
		}
		// Define regular expressions to match the required information
		reBeds := regexp.MustCompile(`(\d+)\s+beds`)
		reBaths := regexp.MustCompile(`(\d+)\s+baths`)
		reSqFt := regexp.MustCompile(`(\d+)\s+Square Feet`)

		// Extract the number of beds, baths, and square footage from the string

		if len(reBeds.FindStringSubmatch(description)) >= 2 && len(reBaths.FindStringSubmatch(description)) >= 2 && len(reSqFt.FindStringSubmatch(description)) >= 2 {
			bedsStr := reBeds.FindStringSubmatch(description)[1]
			bathsStr := reBaths.FindStringSubmatch(description)[1]
			sqFtStr := reSqFt.FindStringSubmatch(description)[1]

			// Convert the extracted strings to integers
			beds, _ := strconv.Atoi(bedsStr)
			baths, _ := strconv.Atoi(bathsStr)
			sqFt, _ := strconv.Atoi(sqFtStr)

			// Format the resulting string
			description = fmt.Sprintf("This property %d Square Feet home has %d baths and %d beds.", sqFt, baths, beds)
		} else {
			description = "Information about property is unavaliable."
		}
	}
	return description
}

func GetZipCode(address string) string {
	re := regexp.MustCompile(`\d{5}(?:[-\s]\d{4})?$`) // Regular expression to match zip code
	matches := re.FindStringSubmatch(address)
	if len(matches) > 0 {
		return matches[0]
	} else {
		return ""
	}
}

func GetUniqueZipCodes(db *gorm.DB) ([]string, error) {
	var properties []Property
	GetAllProperties(db, &properties)

	uniqueZipCodes := make(map[string]bool) // Map to store unique zip codes
	for _, property := range properties {
		zipCode := GetZipCode(property.Address)
		if zipCode != "" {
			uniqueZipCodes[zipCode] = true
		}
	}

	var zipCodeList []string // List to store unique zip codes
	for zipCode := range uniqueZipCodes {
		zipCodeList = append(zipCodeList, zipCode)
	}

	return zipCodeList, nil
}
