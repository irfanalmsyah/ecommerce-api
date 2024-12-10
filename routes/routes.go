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

	protected := app.Group("/", middlewares.Protect)
	protected.Get("/categories/", handlers.GetCategories)
	protected.Get("/products/", handlers.GetProducts)
	protected.Get("/cart/", handlers.GetCart)
	protected.Post("/cart/", middlewares.Validate(&models.AddToCartInput{}), handlers.AddToCart)
	protected.Delete("/cart/", middlewares.Validate(&models.RemoveFromCartInput{}), handlers.RemoveFromCart)
	protected.Post("/checkout/", handlers.Checkout)
}
