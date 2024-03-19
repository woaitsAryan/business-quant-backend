package controllers

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/woaitsAryan/business-quant-backend/helpers"
	"github.com/woaitsAryan/business-quant-backend/initializers"
	"github.com/woaitsAryan/business-quant-backend/models"
)

func GenerateData(c *fiber.Ctx) error {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("No caller information")
	}

	dir := filepath.Dir(filename)

	filePath := filepath.Join(dir, "stock.csv")

	file, err := os.Open(filePath)
	if err != nil {
		helpers.LogServerError("Error opening file", err, "controllers/generate.go")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		helpers.LogServerError("Error reading file", err, "controllers/generate.go")
	}

	records = records[1:]

	tx := initializers.DB.Begin()

	var stocks []models.Stock

	for _, record := range records {
		var gp, fcf, capex, revenue *int64
		date, err := time.Parse("1/2/2006", record[1])
		if err != nil {
			helpers.LogServerError("Error parsing date", err, "controllers/generate.go")
			tx.Rollback()
			os.Exit(1)
		}
		tempRevenue, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			revenue = nil
		}else{
			revenue = &tempRevenue
		}
		tempGp, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			gp = nil
		} else {
			gp = &tempGp
		}
		tempfcf, err := strconv.ParseInt(record[4], 10, 64)
		if err != nil {
			fcf = nil
		} else {
			fcf = &tempfcf
		}
		tempcapex, err := strconv.ParseInt(record[5], 10, 64)
		if err != nil {
			capex = nil
		} else {
			capex = &tempcapex
		}

		stock := models.Stock{
			ID:      uuid.New(),
			Ticker:  record[0],
			Date:    date,
			Revenue: revenue,
			GP:      gp,
			FCF:     fcf,
			Capex:   capex,
		}
		stocks = append(stocks, stock)
	}

	tx.CreateInBatches(stocks, 1000)

	tx.Commit()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Added data successfully!",
	})
}
