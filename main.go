package main

import (
	"log"

	"github.com/irfanalmsyah/ecommerce-api/database"
	"github.com/irfanalmsyah/ecommerce-api/models"
	"github.com/irfanalmsyah/ecommerce-api/routes"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	database.Connect()
	defer database.Close()

	database.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
