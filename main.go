package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/woaitsAryan/business-quant-backend/config"
	"github.com/woaitsAryan/business-quant-backend/helpers"
	"github.com/woaitsAryan/business-quant-backend/initializers"
	"github.com/woaitsAryan/business-quant-backend/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.AddLogger()
	initializers.AutoMigrate()
}

func main() {
	defer initializers.LoggerCleanUp()

	app := fiber.New(fiber.Config{
		ErrorHandler: helpers.ErrorHandler,
		BodyLimit:    config.BODY_LIMIT,
	})

	app.Use(helmet.New())
	app.Use(config.CORS())
	// app.Use(config.RATE_LIMITER())

	app.Use(logger.New())

	app.Static("/", "./public")

	routers.Config(app)

	app.Listen(":" + initializers.CONFIG.PORT)
}
