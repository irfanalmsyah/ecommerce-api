package handlers

import (
	"github.com/irfanalmsyah/ecommerce-api/database"
	"github.com/irfanalmsyah/ecommerce-api/helpers"
	"github.com/irfanalmsyah/ecommerce-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return helpers.InternalServerError(c, "Failed to fetch categories")
	}

	var response []models.CategoryDTO
	for _, category := range categories {
		response = append(response, models.CategoryDTO{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return helpers.ResponseSuccessWithData(c, response)
}

func GetProducts(c *fiber.Ctx) error {
	category := c.Query("category", "")
	var products []models.Product
	if category != "" {
		if err := database.DB.Where("category_id = ?", category).Find(&products).Error; err != nil {
			return helpers.InternalServerError(c, "Failed to fetch products")
		}
	} else {
		if err := database.DB.Find(&products).Error; err != nil {
			return helpers.InternalServerError(c, "Failed to fetch products")
		}
	}

	var response []models.ProductDTO
	for _, product := range products {
		response = append(response, models.ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CategoryID:  product.CategoryID,
		})
	}

	return helpers.ResponseSuccessWithData(c, response)
}
