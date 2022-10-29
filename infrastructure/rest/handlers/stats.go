package handlers

import (
	"github.com/KKrusti/booking/application/useCases"
	"github.com/gofiber/fiber/v2"
)

func CalculateStats(c *fiber.Ctx) error {
	bookingRequest := &[]BookingsRequestDTO{}

	if err := c.BodyParser(bookingRequest); err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return err
		}
	}

	if validateData(c, bookingRequest) {
		return nil
	}
	requestInDomain := mapDtoToDomain(*bookingRequest)
	stats := useCases.CalculateStats(requestInDomain)
	responseDTO := FromDomain(stats)

	err := c.Status(fiber.StatusOK).JSON(responseDTO)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	return nil
}
