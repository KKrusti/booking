package main

import (
	"github.com/KKrusti/booking/infrastructure/rest"
	"github.com/gofiber/fiber"
	"log"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(3000))
}

func setupRoutes(app *fiber.App) {
	rest.StatsControllerEndpoints(app)
	rest.MaximizeControllerEndpoints(app)
}
