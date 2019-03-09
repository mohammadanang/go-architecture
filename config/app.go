package config

import "github.com/jinzhu/gorm"

var ENV map[string]string
var DB *gorm.DB

func init() {
	// Initializing environment project
	ENV = Running()

	// Initializing database connection
	DB = Connect()
}
