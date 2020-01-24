package models

import (

	"github.com/jinzhu/gorm"
	"catalogService/pkg/config"
)

var db *gorm.DB

type Catalog struct {
	gorm.Model

	Product_name string `json:"product_name"`
	Shop_name string `json:"shop_name"`
	Description string `json:"description"`
	Product_type string `json:"product_type"`
	Availability bool `json:"availability"`
	Cost float64 `json:"cost"`

}

func init()  {
	db = config.Connect()
	db.AutoMigrate(&Catalog{})
}

func (catalog *Catalog) CreateProduct() *Catalog  {
	db.NewRecord(catalog)
	db.Create(&catalog)
	return catalog
}

func GetAllProduct() []Catalog {
	var Catlogs []Catalog
	db.Find(&Catlogs)
	return Catlogs
}

func GetProductById(id int64) (*Catalog, *gorm.DB) {
	var catObject Catalog
	db :=db.Where("id = ?",id).Find(&catObject)
	return &catObject,db
}

func DeleteBook(id int64) Catalog  {
	var cat_object Catalog
	db.Where("id = ?",id).Delete(cat_object)
	return cat_object
}