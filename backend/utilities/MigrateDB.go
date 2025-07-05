package utilities

import (
	"fmt"

	"github.com/mission-0/better-stack-backend/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Website{})
	fmt.Println("error while migration", err)
	if err != nil {
		fmt.Println("Not migrated")
		return
	}
}
