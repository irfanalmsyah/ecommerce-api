package routes

import (
	"github.com/irfanalmsyah/ecommerce-api/handlers"
	"github.com/irfanalmsyah/ecommerce-api/middlewares"
	"github.com/irfanalmsyah/ecommerce-api/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register/", middlewares.Validate(&models.RegisterInput{}), handlers.Register)
	app.Post("/login/", middlewares.Validate(&models.LoginInput{}), handlers.Login)
}
