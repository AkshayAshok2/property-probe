package platform

import (
	"PropertyProbe/database"
	"PropertyProbe/platform/properties"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type SuiteTest struct {
	suite.Suite
	db *gorm.DB
}

func TestSuite(t *testing.T) {
	os.Setenv("DB_HOST", "cen3031-project.mysql.database.azure.com")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "go")
	os.Setenv("DB_PASS", "Gators123")
	os.Setenv("DB_DATABASE", "listings")
	defer os.Unsetenv("DB_HOST")
	defer os.Unsetenv("DB_PORT")
	defer os.Unsetenv("DB_USER")
	defer os.Unsetenv("DB_PASS")
	defer os.Unsetenv("DB_DATABASE")

	suite.Run(t, new(SuiteTest))
}

func getModels() []interface{} {
	return []interface{}{&properties.Property{}}
}

// Setup db value
func (t *SuiteTest) SetupSuite() {
	t.db = database.InitDb()

	// Migrate Table
	for _, val := range getModels() {
		t.db.AutoMigrate(val)
	}
}

// Run After All Test Done
func (t *SuiteTest) TearDownSuite() {
	sqlDB, _ := t.db.DB()
	defer sqlDB.Close()

	// Drop Table
	for _, val := range getModels() {
		t.db.Migrator().DropTable(val)
	}
}

// Run Before a Test
func (t *SuiteTest) SetupTest() {

}

// Run After a Test
func (t *SuiteTest) TearDownTest() {

}
