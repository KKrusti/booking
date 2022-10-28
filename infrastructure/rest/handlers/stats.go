package handlers

import (
	"github.com/KKrusti/booking/application/useCases"
	"github.com/KKrusti/booking/domain"
	"github.com/gofiber/fiber"
)

func CalculateStats(c *fiber.Ctx) {
	bookingRequest := &[]domain.Booking{}

	if err := c.BodyParser(bookingRequest); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//TODO add validator?

	stats := useCases.CalculateStats(*bookingRequest)
	responseDTO := FromDomain(stats)

	c.JSON(responseDTO)
}
