package main

import (
	"github.com/KKrusti/booking/infrastructure/rest"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

func SetupRoutes(app *fiber.App) {
	rest.StatsControllerEndpoints(app)
	rest.MaximizeControllerEndpoints(app)
	rest.NotFoundRoute(app)
}
