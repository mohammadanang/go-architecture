package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

/*
	How to use:

	Env["YOUR_KEY"]
*/
func Running() map[string]string {
	result, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error env file")
	}

	return result
}
