package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

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
