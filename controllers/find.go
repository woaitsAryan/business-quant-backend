package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/woaitsAryan/business-quant-backend/initializers"
	"github.com/woaitsAryan/business-quant-backend/utils"
)

func FilterData(c *fiber.Ctx) error {

	db := initializers.DB
	db = utils.GetStockData(db, c)

	fmt.Println(db.Statement.Vars...)

	fmt.Println("Filtering data...")
	return nil
}
