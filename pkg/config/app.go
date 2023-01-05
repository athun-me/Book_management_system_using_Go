package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect(){
	dns := "root:@tcp(127.0.0.1:3306)/bookStore?parseTime=true"
	d, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err !=nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}