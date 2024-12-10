package middlewares

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/irfanalmsyah/ecommerce-api/helpers"
)

var validate = validator.New()

func Validate(data interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodGet {
			return c.Next()
		}

		if err := c.BodyParser(data); err != nil {
			return helpers.BadRequestError(c, "Invalid request body")
		}

		if err := validate.Struct(data); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Field()+" "+err.Tag())
			}
			return helpers.BadRequestError(c, "validation error: " + strings.Join(errors, ", "))
		}

		c.Locals("input", data)

		return c.Next()
	}
}
