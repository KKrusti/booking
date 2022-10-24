package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Post("v1/stats", nil)
	app.Post("v1/maximize", nil)
}
