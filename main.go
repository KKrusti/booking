package main

import (
	"github.com/KKrusti/booking/internal/handlers"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Post("v1/stats", handlers.CalculateStats)
	app.Post("v1/maximize", handlers.Maximize)
}
