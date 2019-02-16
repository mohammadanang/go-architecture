package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

var Env map[string]string

// How to use:
// Env["YOUR_KEY"]
func init() {
	result, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error env file")
	}

	Env = result
}
