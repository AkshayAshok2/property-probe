package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Property struct {
	Date            string  `json:"date"`
	AuctionType     string  `json:"auction_type"`
	JudgementAmount float64 `json:"judgement_amount"`
	Address         string  `json:"address"`
	AssessedValue   float64 `json:"assessedvalue"`
}

/*
func CreatePropertyWithNoConnectionParam(property *Property) (err error) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	CreateProperty(db, property)
}*/

func main() {
	file, err := os.Open("properties.txt")
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
		fmt.Printf("Date: %s\nAuction Type: %s\nJudgementAmount: %.2f\nAddress: %s\nAssessed Value: %.2f\n\n",
			property.Date, property.AuctionType, property.JudgementAmount, property.Address, property.AssessedValue)
	}
}

func parseAmount(amount string) (float64, error) {
	amount = strings.Replace(amount, "$", "", -1)
	amount = strings.Replace(amount, ",", "", -1)
	return strconv.ParseFloat(amount, 64)
}
