package envar

import (
	"log"

	"github.com/joho/godotenv"
)

//loads env. variables from a specified .env file
func LoadEnvironmentals() error {
	err := godotenv.Load("tok.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return err
}
