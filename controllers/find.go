package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/woaitsAryan/business-quant-backend/initializers"
	"github.com/woaitsAryan/business-quant-backend/models"
	"github.com/woaitsAryan/business-quant-backend/utils"
)

func FilterData(c *fiber.Ctx) error {

	db := initializers.DB

	db = utils.GetStockData(db, c)

	var stocks []models.Stock

	db.Find(&stocks)

	return c.Status(200).JSON(stocks)
}
