package rest

import (
	"github.com/KKrusti/booking/infrastructure/rest/handlers"
	"github.com/gofiber/fiber"
)

func StatsControllerEndpoints(app *fiber.App) {
	app.Post("v1/stats", handlers.CalculateStats)
}
