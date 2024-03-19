package initializers

import (
	"fmt"

	"github.com/woaitsAryan/business-quant-backend/models"
)

func AutoMigrate() {
	fmt.Println("\nStarting Migrations...")
	DB.AutoMigrate(
		&models.Stock{},

	)
	fmt.Println("Migrations Finished!")
}
