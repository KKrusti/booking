package handlers

import (
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"github.com/KKrusti/booking/internal/core/services"
	"github.com/gofiber/fiber"
)

func CalculateStats(c *fiber.Ctx) {
	bookingRequest := &[]entities.Booking{}

	if err := c.BodyParser(bookingRequest); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//TODO add validator?

	stats := services.CalcStats(*bookingRequest)
	c.JSON(stats.ToStatsResponse())
}
