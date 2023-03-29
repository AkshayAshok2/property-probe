package properties

import (
	"gorm.io/gorm"
)

type Property struct {
	Date            string  `json:"date"`
	AuctionType     string  `json:"auction_type"`
	JudgementAmount float64 `json:"judgement_amount"`
	Address         string  `json:"address"`
	AssessedValue   float64 `json:"assessedvalue"`
	LatLon          string  `json:"latlon"`
}

func CreateProperty(db *gorm.DB, property *Property) (err error) {
	err = db.Create(property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProperties(db *gorm.DB, property *[]Property) (err error) {
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

func UpdateProperty(db *gorm.DB, property *Property) (err error) {
	db.Save(property)
	return nil
}

func DeleteProperty(db *gorm.DB, property *Property, address string) (err error) {
	db.Where("address = ?", address).Delete(property)
	return nil
}
