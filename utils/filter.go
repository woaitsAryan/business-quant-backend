package utils

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/woaitsAryan/business-quant-backend/helpers"
	"gorm.io/gorm"
	"strings"
)

func GetStockData(db *gorm.DB, c *fiber.Ctx) *gorm.DB {
	ticker := c.Query("ticker")
	column := c.Query("column")
	period := c.Query("period")

	columns := strings.Split(column, ",")
	query := db.Select(columns).Where("ticker = ?", ticker)

	if period != "" {
		yearsStr := strings.TrimSuffix(period, "y")
		years, err := strconv.Atoi(yearsStr)
		if err != nil {
			helpers.LogServerError("Error parsing period", err, "utils/filter.go")
			return nil
		}
		Date := time.Now().AddDate(-years, 0, 0)
		query = query.Where("date >= ?", Date)
	}
	return query
}
