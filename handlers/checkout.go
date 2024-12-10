package handlers

import (
	"github.com/irfanalmsyah/ecommerce-api/database"
	"github.com/irfanalmsyah/ecommerce-api/helpers"
	"github.com/irfanalmsyah/ecommerce-api/models"

	"github.com/gofiber/fiber/v2"
)

func Checkout(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var cart models.Cart
	if err := database.DB.Preload("CartItems.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return helpers.NotFoundError(c, "Cart not found")
	}

	if len(cart.CartItems) == 0 {
		return helpers.BadRequestError(c, "Cart is empty")
	}

	var total float64
	for _, item := range cart.CartItems {
		total += item.Product.Price * float64(item.Quantity)
	}

	// payment processing here

	order := models.Order{
		UserID: userID,
		Total:  total,
		Status: "Completed",
	}
	database.DB.Create(&order)

	for _, item := range cart.CartItems {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}
		database.DB.Create(&orderItem)
	}

	database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})

	return helpers.ResponseSuccessWithData(c, fiber.Map{"order_id": order.ID})
}
