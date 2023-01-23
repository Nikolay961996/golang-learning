package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

/*
Your account number is: 738556
Host: sql12.freesqldatabase.com
Database name: sql12592850
Database user: sql12592850
Database password: WtQ1LBtdUZ
Port number: 3306
*/

func Connect() {
	d, err := gorm.Open("mysql", "sql12592850:WtQ1LBtdUZ@tcp(sql12.freesqldatabase.com:3306)/sql12592850?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}