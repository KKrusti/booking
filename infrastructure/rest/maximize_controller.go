package rest

import (
	"github.com/KKrusti/booking/infrastructure/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func MaximizeControllerEndpoints(app *fiber.App) {
	app.Post("v1/maximize", handlers.Maximize)
}
