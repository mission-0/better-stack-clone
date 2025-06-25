package utilities

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDb() {
	db_url := os.Getenv("DATABASE_URL")

	if db_url == "" {
		fmt.Print("empty url")
	}

	fmt.Println("env ", string(db_url))
	fmt.Println("env ", string(os.Getenv("JWT_SECRET")))

	DB, _ = gorm.Open(postgres.Open(db_url), &gorm.Config{})

	fmt.Print("Db", DB)
	if err != nil {
		println("db connection error")
	}

	fmt.Println("db connected")
	if err != nil {
		log.Fatal("something went wrong")
	}

	// Other DB Operations like CRUD ahead
}
