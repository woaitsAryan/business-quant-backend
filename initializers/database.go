package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	DB_PASSWORD := CONFIG.DB_PASSWORD

	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/stock", DB_PASSWORD)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
	})

	if err != nil {
		log.Fatal("Failed to Connect to the database")
	} else {
		log.Println("Connected to database!")
	}
}
