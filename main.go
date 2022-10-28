package main

import (
	"github.com/KKrusti/booking/infrastructure/rest"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	rest.StatsControllerEndpoints(app)
	rest.MaximizeControllerEndpoints(app)
}
