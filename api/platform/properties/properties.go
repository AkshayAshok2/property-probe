package properties

import (
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	AuctionType     string `json:"auction_type"`
	JudgementAmount string `json:"judgement_amount"`
	Address         string `json:"address"`
	AssessedValue   string `json:"assessedvalue"`
}

func CreateProperty(db *gorm.DB, Property *Property) (err error) {
	err = db.Create(Property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProperties(db *gorm.DB, Property *[]Property) (err error) {
	err = db.Find(Property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProperty(db *gorm.DB, Property *Property, address string) (err error) {
	err = db.Where("address = ?", address).First(Property).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProperty(db *gorm.DB, Property *Property) (err error) {
	db.Save(Property)
	return nil
}

func DeleteProperty(db *gorm.DB, Property *Property) (err error) {
	db.Delete(Property)
	return nil
}

// type Getter interface {
// 	GetAll() []Property
// }

// type Adder interface {
// 	Add(property Property)
// }

// type Repo struct {
// 	Properties []Property
// }

// func New() *Repo {
// 	return &Repo{
// 		Properties: []Property{},
// 	}
// }

// func (r *Repo) Add(property Property) {
// 	r.Properties = append(r.Properties, property)
// }

// func (r *Repo) GetAll() []Property {
// 	return r.Properties
// }
