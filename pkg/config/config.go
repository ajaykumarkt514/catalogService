package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect()  *gorm.DB{
	var db *gorm.DB
	d,err := gorm.Open("mysql","root:123456780@/catalog?parseTime=true")
	if err!=nil{
		log.Fatal(err)
		panic("Unable to Connect to database, checkit!")
	}
	db = d
	return db

}
