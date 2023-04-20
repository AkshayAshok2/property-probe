package properties

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateProperty(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Exec("DELETE FROM properties").Error
	assert.NoError(t, err)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}
	err = CreateProperty(db, &prop)
	assert.NoError(t, err)
	assert.NotZero(t, prop.AssessedValue)

	DeleteProperty(db, &prop, prop.Address)
}

func TestGetAllProperties(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Exec("DELETE FROM properties").Error
	assert.NoError(t, err)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street 12345",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}
	CreateProperty(db, &prop)

	// Get all properties
	var props []Property
	err = GetAllProperties(db, &props)
	assert.NoError(t, err)
	assert.Len(t, props, 1)

	DeleteProperty(db, &prop, prop.Address)
}

func TestUpdateProperty(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Exec("DELETE FROM properties").Error
	assert.NoError(t, err)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}
	CreateProperty(db, &prop)
	prop.JudgementAmount = 55555.55
	err = UpdateProperty(db, &prop)
	assert.NoError(t, err)

	DeleteProperty(db, &prop, "123 Main Street")
}

func TestDeleteProperty(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Exec("DELETE FROM properties").Error
	assert.NoError(t, err)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}

	CreateProperty(db, &prop)

	// Delete a property
	delete := &Property{}
	err = DeleteProperty(db, delete, prop.Address)
	assert.NoError(t, err)

	// Get all properties again (should be empty)
	var props []Property
	err = GetAllProperties(db, &props)
	assert.NoError(t, err)
	assert.Len(t, props, 0)
}

func TestGetProperty(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Exec("DELETE FROM properties").Error
	assert.NoError(t, err)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	property := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}

	CreateProperty(db, &property)

	found := &Property{}
	err = GetProperty(db, found, "123 Main Street")
	assert.NoError(t, err)
	assert.Equal(t, found.Address, "123 Main Street")

	DeleteProperty(db, &property, property.Address)
}

// // Unreliable test based on scraping issues due to Google rearranging search results
// func TestGetDescription(t *testing.T) {
// 	description := GetDescription("1013 Fieldstone Drive, 32940")
// 	assert.Equal(t, description, "This property 1921 Square Feet home has 3 baths and 2 beds.")
// }

func TestGetZipCodeProperties(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}
	CreateProperty(db, &prop)

	prop = Property{
		AuctionType:     "public",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}
	CreateProperty(db, &prop)

	// Get zipcode properties
	var props []Property
	err = GetZipCodeProperties(db, &props, "32940")
	assert.NoError(t, err)
	assert.Len(t, props, 2)

	DeleteProperty(db, &prop, prop.Address)
}

func TestGetZipCode(t *testing.T) {
	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "20308 NW COUNTY RD 2054 ALACHUA, FL 32615",
		AssessedValue:   9876.54,
		ZipCode:         "",
		Description:     "2400 sqft, 2 bed 3 bath",
	}

	zip := GetZipCode(prop.Address)
	assert.Equal(t, zip, "32615")

}

func TestGetUnqiueZipCodes(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	err = db.Exec("DELETE FROM properties").Error
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}
	CreateProperty(db, &prop)

	prop = Property{
		AuctionType:     "public",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
		ZipCode:         "32940",
		Description:     "2400 sqft, 2 bed 3 bath",
	}

	// Get zipcodes
	zipcodes, err := GetUniqueZipCodes(db)

	assert.NoError(t, err)
	assert.Len(t, zipcodes, 1)
}
