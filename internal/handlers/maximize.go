package handlers

import (
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"github.com/KKrusti/booking/internal/core/services"
	"github.com/gofiber/fiber"
)

func Maximize(c *fiber.Ctx) {
	bookingRequest := &[]entities.Booking{}

	if err := c.BodyParser(bookingRequest); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//TODO add validator?
	statsCalculation := services.Maximize(*bookingRequest)
	statsResponse := statsCalculation.ToMaxResponse()

	c.JSON(statsResponse)
}
