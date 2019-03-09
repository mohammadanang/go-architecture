package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbConfig struct {
	name     string
	username string
	password string
	host     string
	port     string
}

func Connect() *gorm.DB {
	var result *gorm.DB
	dialect := ENV["DB_DIALECT"]
	setup := dbConfig{
		name:     ENV["DB_NAME"],
		username: ENV["DB_USERNAME"],
		password: ENV["DB_PASSWORD"],
		host:     ENV["DB_HOST"],
		port:     ENV["DB_PORT"],
	}

	if dialect == "mysql" {
		result = mySQLConnect(setup)
	} else {
		result = postgresConnect(setup)
	}

	return result
}

func mySQLConnect(s dbConfig) *gorm.DB {
	port := s.port
	if port == "" {
		port = "3306"
	}

	result, err := gorm.Open("mysql", s.username+":"+s.password+"@tcp("+s.host+":"+port+")/"+s.name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Database connection error", err)
	}

	return result
}

func postgresConnect(s dbConfig) *gorm.DB {
	port := s.port
	if port == "" {
		port = "5432"
	}

	result, err := gorm.Open("postgres", "host="+s.host+" port="+port+" user="+s.username+" dbname="+s.name+" password="+s.password)
	if err != nil {
		fmt.Println("Database connection error", err)
	}

	return result
}
