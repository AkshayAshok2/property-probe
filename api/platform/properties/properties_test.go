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

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main St",
		AssessedValue:   9876.54,
	}
	err = CreateProperty(db, &prop)
	assert.NoError(t, err)
	assert.NotZero(t, prop.AssessedValue)

	DeleteProperty(db, &prop, prop.Address)
}

func TestGetProperties(t *testing.T) {
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
		Address:         "123 Main St",
		AssessedValue:   9876.54,
	}
	CreateProperty(db, &prop)

	// Get all properties
	var props []Property
	err = GetProperties(db, &props)
	assert.NoError(t, err)
	assert.Len(t, props, 1)

	DeleteProperty(db, &prop, prop.Address)
}

func TestUpdatePro(t *testing.T) {
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
		Address:         "123 Main St",
		AssessedValue:   9876.54,
	}
	CreateProperty(db, &prop)
	prop.JudgementAmount = 55555.55
	err = UpdateProperty(db, &prop)
	assert.NoError(t, err)

	DeleteProperty(db, &prop, "123 Main St")
}

func TestDeleteProperty(t *testing.T) {
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
		Address:         "123 Main St",
		AssessedValue:   9876.54,
	}

	CreateProperty(db, &prop)

	// Delete a property
	err = DeleteProperty(db, &prop, prop.Address)
	assert.NoError(t, err)

	// Get all properties again (should be empty)
	var props []Property
	err = GetProperties(db, &props)
	assert.NoError(t, err)
	assert.Len(t, props, 0)
}

// func TestGetProperty(t *testing.T) {
// 	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	assert.NoError(t, err)
// 	assert.NotNil(t, db)

// 	// Migrate the database schema
// 	err = db.AutoMigrate(Property{})
// 	assert.NoError(t, err)

// 	property := Property{
// 		AuctionType:     "private",
// 		JudgementAmount: 12345.67,
// 		Address:         "123 Main St",
// 		AssessedValue:   9876.54,
// 	}
// 	err = CreateProperty(db, &property)

// 	propTest := Property{
// 		AuctionType:     "test",
// 		JudgementAmount: 0,
// 		Address:         "test",
// 		AssessedValue:   0,
// 	}
// 	err = GetProperty(db, &propTest, "123 Main Street")
// 	assert.NoError(t, err)
// 	assert.Equal(t, propTest.Address, "123 Main Street")
// }

// func TestPropertyCRUD(t *testing.T) {
// 	// Connect to an in-memory SQLite database for testing
// 	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	assert.NoError(t, err)
// 	assert.NotNil(t, db)

// 	// Migrate the database schema
// 	err = db.AutoMigrate(Property{})
// 	assert.NoError(t, err)

// 	// Create a property
// 	prop := &Property{
// 		auction_type:     "public",
// 		judgement_amount: 12345.67,
// 		address:          "123 Main St",
// 		assessed_value:   9876.54,
// 	}
// 	err = CreateProperty(db, prop)
// 	assert.NoError(t, err)
// 	assert.NotZero(t, prop.ID)

// 	// Get all properties
// 	var props []Property
// 	err = GetProperties(db, &props)
// 	assert.NoError(t, err)
// 	assert.Len(t, props, 1)

// 	// Get a single property by address
// 	var prop2 Property
// 	err = GetProperty(db, &prop2, prop.address)
// 	assert.NoError(t, err)
// 	assert.Equal(t, prop.ID, prop2.ID)

// 	// Update a property
// 	prop.judgement_amount = 55555.55
// 	err = UpdateProperty(db, prop)
// 	assert.NoError(t, err)

// 	// Get the updated property
// 	var prop3 Property
// 	err = GetProperty(db, &prop3, prop.address)
// 	assert.NoError(t, err)
// 	assert.Equal(t, prop.judgement_amount, prop3.judgement_amount)

// 	// Delete a property
// 	err = DeleteProperty(db, prop, prop.address)
// 	assert.NoError(t, err)

// 	// Get all properties again (should be empty)
// 	err = GetProperties(db, &props)
// 	assert.NoError(t, err)
// 	assert.Len(t, props, 0)
// }

// func Add(property Property) {
// 	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	result := db.Create(property)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}

// }
// func getSize() int {
// 	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	var properties []Property
// 	resultCheckFirst := db.Find(&properties)
// 	if resultCheckFirst.Error != nil {
// 		panic(resultCheckFirst.Error)
// 	}
// 	return len(properties)
// }

// func TestAdd(t *testing.T) {
// 	//Size before add function

// 	var sizeInitial = getSize()

// 	//Creates example property
// 	property := Property{
// 		auction_type:     "auction_type_1",
// 		judgement_amount: 1000.0,
// 		address:          "123 Main St",
// 		assessed_value:   5000.0,
// 	}

// 	// Insert the property into the database
// 	Add(property)

// 	var sizeAfter = getSize()

// 	if sizeAfter <= sizeInitial {
// 		t.Errorf("No property was added")
// 	} else {
// 		fmt.Println("Test passed: property added")
// 	}

// }
