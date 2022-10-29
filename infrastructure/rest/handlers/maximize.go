package handlers

import (
	"fmt"
	"github.com/KKrusti/booking/application/useCases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

func Maximize(c *fiber.Ctx) {
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
	statsCalculation := useCases.Maximize(requestInDomain)
	err := c.Status(fiber.StatusOK).JSON(statsCalculation)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return
	}
}

func validateData(c *fiber.Ctx, bookingRequest *[]BookingsRequestDTO) bool {
	validate := validator.New()
	var err error
	for _, element := range *bookingRequest {
		err = validate.Struct(element)
		if err != nil {
			errorMsg := fmt.Sprintf("There's a problem with the input data %s", err.Error())
			err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   errorMsg,
			})
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				return false
			}
			return true
		}
	}
	return false
}
