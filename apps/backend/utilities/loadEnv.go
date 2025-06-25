package utilities

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVaribales() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no env variable")
	}
}
