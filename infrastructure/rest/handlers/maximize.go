package handlers

import (
	"github.com/KKrusti/booking/application/useCases"
	"github.com/KKrusti/booking/domain"
	"github.com/gofiber/fiber"
)

func Maximize(c *fiber.Ctx) {
	bookingRequest := &[]domain.Booking{}

	if err := c.BodyParser(bookingRequest); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//TODO add validator?
	statsCalculation := useCases.Maximize(*bookingRequest)

	c.JSON(statsCalculation)
}
