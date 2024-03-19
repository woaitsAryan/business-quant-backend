package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/woaitsAryan/business-quant-backend/controllers"
)

func FilterRouter(app *fiber.App) {
	app.Get("/", controllers.FilterData)
}
