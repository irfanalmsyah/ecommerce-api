package helpers

import (
	"github.com/irfanalmsyah/ecommerce-api/models"
	"github.com/gofiber/fiber/v2"
)

func ResponseSuccess(c *fiber.Ctx, message string) error {
	return jsonResponse(c, fiber.StatusOK, message, nil)
}

func ResponseSuccessWithData(c *fiber.Ctx, data interface{}) error {
	return jsonResponse(c, fiber.StatusOK, "success", data)
}

func ResponseCreated(c *fiber.Ctx, message string) error {
	return jsonResponse(c, fiber.StatusCreated, message, nil)
}

func ResponseCreatedWithData(c *fiber.Ctx, message string, data interface{}) error {
	return jsonResponse(c, fiber.StatusCreated, message, data)
}

func BadRequestError(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusBadRequest, message)
}

func InternalServerError(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusInternalServerError, message)
}

func ConflictError(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusConflict, message)
}

func UnauthorizedError(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusUnauthorized, message)
}

func NotFoundError(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusNotFound, message)
}

func jsonResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := models.Response{
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}

func responseError(c *fiber.Ctx, statusCode int, message string) error {
	return jsonResponse(c, statusCode, message, nil)
}
