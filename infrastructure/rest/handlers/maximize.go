package handlers

import (
	"github.com/KKrusti/booking/application/useCases"
	"github.com/gofiber/fiber"
)

func Maximize(c *fiber.Ctx) {
	bookingRequest := &[]BookingsRequestDTO{}

	if err := c.BodyParser(bookingRequest); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//TODO add validator?
	requestInDomain := mapDtoToDomain(*bookingRequest)
	statsCalculation := useCases.Maximize(requestInDomain)

	c.JSON(statsCalculation)
}
