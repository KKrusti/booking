package handlers

import (
	"github.com/KKrusti/booking/application/useCases"
	"github.com/gofiber/fiber"
)

func CalculateStats(c *fiber.Ctx) {
	bookingRequest := &[]BookingsRequestDTO{}

	if err := c.BodyParser(bookingRequest); err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return
		}
	}

	if validateData(c, bookingRequest) {
		return
	}
	requestInDomain := mapDtoToDomain(*bookingRequest)
	stats := useCases.CalculateStats(requestInDomain)
	responseDTO := FromDomain(stats)

	err := c.Status(fiber.StatusOK).JSON(responseDTO)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return
	}
}
