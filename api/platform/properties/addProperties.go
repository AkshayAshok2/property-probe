package properties

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreatePropertyWithNoConnectionParam(property *Property) (err error) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	defer db.Close()
	CreateProperty(db, property)
	return err
}

func parseAmount(amount string) (float64, error) {
	amount = strings.Replace(amount, "$", "", -1)
	amount = strings.Replace(amount, ",", "", -1)
	return strconv.ParseFloat(amount, 64)
}

func AddPropertiesToDatbase() {
	// get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// specify the relative path to your file
	relPath := "/platform/properties/properties.txt"

	// join the current working directory with the file path
	absPath := filepath.Join(cwd, relPath)

	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var properties []Property
	var currentProperty *Property
	var date string = ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "0") || strings.HasPrefix(line, "1") {
			if currentProperty != nil {
				currentProperty.Date = date
				fmt.Println("Date when adding to vector: " + currentProperty.Date)
				fmt.Println("")
				properties = append(properties, *currentProperty)
			}
			date = strings.TrimSpace(line)
		} else if strings.HasPrefix(line, "[") {
			currentProperty.LatLon = strings.TrimSpace(line)
		} else {
			fields := strings.Split(line, ": ")

			if len(fields) == 2 {
				switch fields[0] {
				case "Auction Type":
					auctionType := fields[1]
					fmt.Println("Adding Auction Type:", auctionType)
					currentProperty = &Property{AuctionType: auctionType}
				case "Final Judgment Amount":
					judgementAmount, err := parseAmount(fields[1])
					if err != nil {
						fmt.Println("Error parsing judgement amount:", err)
					} else {
						fmt.Println("Adding Judgement Amount:", judgementAmount)
						currentProperty.JudgementAmount = judgementAmount
					}
				case "Property Address":
					address := fields[1]
					fmt.Println("Adding Address:", address)
					currentProperty.Address = address
				case "Assessed Value":
					assessedValue, err := parseAmount(fields[1])
					if err != nil {
						fmt.Println("Error parsing assessed value:", err)
					} else {
						fmt.Println("Adding Assessed Value:", assessedValue)
						currentProperty.AssessedValue = assessedValue
					}
				}
			}
		}

	}
	if currentProperty != nil {
		currentProperty.Date = date
		properties = append(properties, *currentProperty)

	}

	for _, property := range properties {
		property.Description = GetDescription(property.Address)
		property.ZipCode = GetZipCode(property.Address)
		if property.Address != "" {
			CreatePropertyWithNoConnectionParam(&property)
		}
	}
}
