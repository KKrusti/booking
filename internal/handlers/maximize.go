package handlers

import (
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"github.com/KKrusti/booking/internal/core/services"
	"github.com/gofiber/fiber"
)

func Maximize(c *fiber.Ctx) {
	bookingRequest := &[]entities.Request{}

	if err := c.BodyParser(bookingRequest); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//TODO add validator?
	services.Maximize(*bookingRequest)

	//c.JSON(statsResponse)
}
