package rest

import (
	"github.com/KKrusti/booking/infrastructure/rest/handlers"
	"github.com/gofiber/fiber"
)

func MaximizeControllerEndpoints(app *fiber.App) {
	app.Post("v1/maximize", handlers.Maximize)
}
