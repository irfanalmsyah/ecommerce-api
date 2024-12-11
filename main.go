package main

import (
	"log"

	"github.com/irfanalmsyah/ecommerce-api/database"
	"github.com/irfanalmsyah/ecommerce-api/models"
	"github.com/irfanalmsyah/ecommerce-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	defer database.Close()

	database.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

	if database.IsTableEmpty(&models.Category{}) {
		if err := database.SeedDatabase("seed.sql"); err != nil {
			log.Printf("Failed to seed database: %v\n", err)
		}
	}

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
