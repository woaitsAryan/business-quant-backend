package routers

import (
	"github.com/gofiber/fiber/v2"
)

func Config(app *fiber.App) {
	GenerateRouter(app)
	FilterRouter(app)
}
