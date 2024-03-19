package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	// "github.com/woaitsAryan/business-quant-backend/helpers"
	"gorm.io/gorm"
	// "strconv"
	"strings"
	// "time"
)

func GetStockData(db *gorm.DB, c *fiber.Ctx) *gorm.DB {
	ticker := c.Query("ticker")
	column := c.Query("column")
	// period := c.Query("period")
	columns := strings.Split(column, ",")
	fmt.Println(columns)
	query := db.Debug().Select(columns).Where("ticker = ?", ticker)
	fmt.Println(query)
	// if period != "" {
	// 	years, err := strconv.Atoi(period)
	// 	if err != nil {
	// 		helpers.LogServerError("Error parsing period", err, "utils/filter.go")
	// 	}
	// 	startDate := time.Now().AddDate(-years, 0, 0)
	// 	query = query.Where("date >= ?", startDate)
	// }

	return query
}
