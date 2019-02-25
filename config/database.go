package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	result, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Database connection error")
	}
	defer result.Close()

	DB = result
}
