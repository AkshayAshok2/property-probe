func TestAdd(t *testing.T) {
	fmt.Println("TESTING ADD")
	dsn := "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var properties []Property
	resultCheckFirst := db.Find(&properties)
	if resultCheckFirst.Error != nil {
		panic(resultCheckFirst.Error)
	}
	//Size before add function
	var sizeInitial = len(properties)

	property := Property{
		auction_type:     "auction_type_1",
		judgement_amount: 1000.0,
		address:          "123 Main St",
		assessed_value:   5000.0,
	}

	// Insert the property into the database
	result := db.Create(&property)
	if result.Error != nil {
		panic(result.Error)
	}
	resultCheck := db.Find(&properties)
	if resultCheck.Error != nil {
		panic(resultCheck.Error)
	}
	//Checks the size after
	var sizeAfter = len(properties)

	if sizeAfter <= sizeInitial {
		t.Errorf("No property was added")
	} else {
		fmt.Println("Test passed")
	}

}

type Property struct {
	auction_type     string
	judgement_amount float64
	address          string
	assessed_value   float64
}
