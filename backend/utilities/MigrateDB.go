package utilities

import (
	"fmt"

	"github.com/mission-0/better-stack-backend/models"
)

func MigrateDB() {
	// DB.Migrator().DropTable(&models.Logs{})
	// DB.Migrator().DropTable(&models.Website{})
	// DB.Migrator().DropTable(&models.User{})
	err := DB.AutoMigrate(&models.User{}, &models.Website{}, &models.Logs{})
	fmt.Println("error while migration", err)
	if err != nil {
		fmt.Println("Not migrated")
		return
	}
}
