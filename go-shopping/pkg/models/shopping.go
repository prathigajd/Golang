package models

import (
	"github.com/jinzhu/gorm"
	"github.com/prathigajd/go-shopping/pkg/config"
)

var db *gorm.DB

type Shopping struct{
	gorm.Model
    Snumber string  `json:"Snumber"`
	Items string     `json:"Items"`
	Price string      `json:"Price"`

}
func init() {
config.Connect()
db = config.GetDB()
db.AutoMigrate(&Shopping{})
}

func (s *Shopping) Createshopping() *Shopping{
  db.NewRecord(s) 
  db.Create(&s)
  return s
}

func GetAllShoppings() []Shopping {
	var Shoppings [] Shopping
	db.Find(&Shoppings)
	return Shoppings
}

func GetShoppingByID(Id int64) (*Shopping, *gorm.DB){
	var getShopping Shopping 
	db := db.Where("ID=?", Id).Find(&getShopping)
	return &getShopping, db

}

func DeleteShopping(ID int64) Shopping{
	var Shopping Shopping 
	db.Where("ID=?", ID).Delete(Shopping)
	return Shopping
}
