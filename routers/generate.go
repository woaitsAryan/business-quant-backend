package routers

import (
	"github.com/woaitsAryan/business-quant-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func GenerateRouter(app *fiber.App) {
	app.Get("/generate", controllers.GenerateData)
}
