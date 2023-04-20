package handler

import (
	"PropertyProbe/database"
	"PropertyProbe/platform/properties"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProperty(t *testing.T) {
	// Initialize a new PropertyRepo
	repo := New()

	router := gin.Default()
	router.POST("/property", repo.CreateProperty)

	// Define test case
	requestBody := `{
				"auction_type": "foreclosure",
				"judgement_amount": 50000.0,
				"address": "123 Main St",
				"assessedvalue": 75000.0,
				"description": "This is a test property",
				"zipcode": "12345"
				
			}`
	req, err := http.NewRequest("POST", "/property", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call CreateProperty handler function
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse the response body
	var property properties.Property
	err = json.NewDecoder(rr.Body).Decode(&property)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Validate the created property fields
	assert.Equal(t, "foreclosure", property.AuctionType)
	assert.Equal(t, 50000.0, property.JudgementAmount)
	assert.Equal(t, "123 Main St", property.Address)
	assert.Equal(t, 75000.0, property.AssessedValue)
	assert.Equal(t, "This is a test property", property.Description)
	assert.Equal(t, "", property.ZipCode)

	database.ClearDB()
}

func TestGetAllProperties(t *testing.T) {
	// Initialize a new router and repository
	r := gin.Default()
	repo := New()

	// Add a few properties to the database
	properties_ := []properties.Property{
		{AuctionType: "Foreclosure", JudgementAmount: 100000, Address: "123 Main St", AssessedValue: 80000, Description: "Nice house", ZipCode: "12345"},
		{AuctionType: "Tax Sale", JudgementAmount: 50000, Address: "456 Elm St", AssessedValue: 40000, Description: "Cozy cottage", ZipCode: "67890"},
		{AuctionType: "Foreclosure", JudgementAmount: 75000, Address: "789 Oak Ave", AssessedValue: 60000, Description: "Spacious ranch", ZipCode: "54321"},
	}
	for _, p := range properties_ {
		err := repo.Db.Create(&p).Error
		if err != nil {
			t.Fatalf("Failed to create property: %v", err)
		}
	}

	// Create a new HTTP request to get all properties
	req, err := http.NewRequest("GET", "/properties", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Call GetAllProperties with the request and recorder
	r.GET("/properties", repo.GetAllProperties)
	r.ServeHTTP(rr, req)

	// Check that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	//Parse the response body into a slice of properties
	var response []properties.Property
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	//Check that the response matches the properties in the database
	assert.Equal(t, len(properties_), len(response))
	for i, p := range properties_ {
		assert.Equal(t, p.AuctionType, response[i].AuctionType)
		assert.Equal(t, p.JudgementAmount, response[i].JudgementAmount)
		assert.Equal(t, p.Address, response[i].Address)
		assert.Equal(t, p.AssessedValue, response[i].AssessedValue)
		assert.Equal(t, p.Description, response[i].Description)
		assert.Equal(t, p.ZipCode, response[i].ZipCode)
	}
	database.ClearDB()
}

func TestGetZipCodeProperties(t *testing.T) {
	// Initialize a new router and repository
	r := gin.Default()
	repo := New()

	// Add a few properties to the database
	properties_ := []properties.Property{
		{AuctionType: "type1", JudgementAmount: 100, Address: "addr1", AssessedValue: 200, Description: "desc1", ZipCode: "12345"},
		{AuctionType: "type2", JudgementAmount: 200, Address: "addr2", AssessedValue: 300, Description: "desc2", ZipCode: "23456"},
		{AuctionType: "type3", JudgementAmount: 300, Address: "addr3", AssessedValue: 400, Description: "desc3", ZipCode: "12345"},
	}
	for _, p := range properties_ {
		err := repo.Db.Create(&p).Error
		if err != nil {
			t.Fatalf("Failed to create property: %v", err)
		}
	}

	// Create a new HTTP request to get properties with zip code "12345"
	req, err := http.NewRequest("GET", "/properties/12345", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Call GetAllProperties with the request and recorder
	r.GET("/properties/:zipcode", repo.GetZipCodeProperties)
	r.ServeHTTP(rr, req)

	// Check that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	//Parse the response body into a slice of properties
	var response []properties.Property
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check that we only got properties with zip code "11111"
	for _, prop := range response {
		if prop.ZipCode != "12345" {
			t.Errorf("Expected property with zip code 12345 but got zip code %s", prop.ZipCode)
		}
	}
}

func TestGetProperty(t *testing.T) {
	// Initialize a new router and repository
	r := gin.Default()
	repo := New()

	// Add a few properties to the database
	properties_ := []properties.Property{
		{AuctionType: "type1", JudgementAmount: 100, Address: "122 Main St", AssessedValue: 200, Description: "desc1", ZipCode: "12345"},
		{AuctionType: "type2", JudgementAmount: 200, Address: "456 Elm St", AssessedValue: 300, Description: "desc2", ZipCode: "23456"},
		{AuctionType: "type3", JudgementAmount: 300, Address: "123 Main St", AssessedValue: 400, Description: "desc3", ZipCode: "12345"},
	}
	for _, p := range properties_ {
		err := repo.Db.Create(&p).Error
		if err != nil {
			t.Fatalf("Failed to create property: %v", err)
		}
	}

	// Create a new HTTP request to get properties by address
	address := "123 Main St"
	req, err := http.NewRequest("GET", "/properties/address/"+address, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Call GetAddressProperties with the request and recorder
	r.GET("/properties/address/:address", repo.GetProperty)
	r.ServeHTTP(rr, req)

	// Check that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	var response properties.Property
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if response.Address != "123 Main St" {
		t.Errorf("Expected property with address 123 Main St but got address %s", response.Address)
	}

	database.ClearDB()
}

func TestGetUniqueZipCodes(t *testing.T) {
	// Initialize a new router and repository
	r := gin.Default()
	repo := New()

	// Add a few properties to the database
	properties_ := []properties.Property{
		{AuctionType: "type1", JudgementAmount: 100, Address: "addr1", AssessedValue: 200, Description: "desc1", ZipCode: "12345"},
		{AuctionType: "type2", JudgementAmount: 200, Address: "addr2", AssessedValue: 300, Description: "desc2", ZipCode: "23456"},
		{AuctionType: "type3", JudgementAmount: 300, Address: "addr3", AssessedValue: 400, Description: "desc3", ZipCode: "12345"},
	}
	for _, p := range properties_ {
		err := repo.Db.Create(&p).Error
		if err != nil {
			t.Fatalf("Failed to create property: %v", err)
		}
	}

	// Create a new HTTP request to get properties with zip code "12345"
	req, err := http.NewRequest("GET", "/properties/zipcodes", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Call GetAllProperties with the request and recorder
	r.GET("/properties/zipcodes", repo.GetUniqueZipCodes)
	r.ServeHTTP(rr, req)

	// Check that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	//Parse the response body into a slice of properties
	var response []string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check that we only got 2 unique zip codes
	assert.Equal(t, len(response), 2)

}
