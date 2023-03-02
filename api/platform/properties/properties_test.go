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
		Address:         "123 Main Street",
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
		Address:         "123 Main Street",
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

func TestUpdateProperty(t *testing.T) {
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

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	prop := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
	}

	CreateProperty(db, &prop)

	// Delete a property
	delete := &Property{}
	err = DeleteProperty(db, delete, prop.Address)
	assert.NoError(t, err)

	// Get all properties again (should be empty)
	var props []Property
	err = GetProperties(db, &props)
	assert.NoError(t, err)
	assert.Len(t, props, 0)
}

func TestGetProperty(t *testing.T) {
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Migrate the database schema
	err = db.AutoMigrate(Property{})
	assert.NoError(t, err)

	property := Property{
		AuctionType:     "private",
		JudgementAmount: 12345.67,
		Address:         "123 Main Street",
		AssessedValue:   9876.54,
	}

	CreateProperty(db, &property)

	found := &Property{}
	err = GetProperty(db, found, "123 Main Street")
	assert.NoError(t, err)
	assert.Equal(t, found.Address, "123 Main Street")

	DeleteProperty(db, &property, property.Address)
}
