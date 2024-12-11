package handlers

import (
	"github.com/irfanalmsyah/ecommerce-api/database"
	"github.com/irfanalmsyah/ecommerce-api/helpers"
	"github.com/irfanalmsyah/ecommerce-api/models"

	"github.com/gofiber/fiber/v2"
)

func AddToCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input := c.Locals("input").(*models.AddToCartInput)

	var product models.Product
	if err := database.DB.First(&product, input.ProductID).Error; err != nil {
		return helpers.NotFoundError(c, "Product not found")
	}

	var cart models.Cart
	if err := database.DB.Preload("CartItems").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		cart = models.Cart{UserID: userID}
		database.DB.Create(&cart)
	}

	var cartItem models.CartItem
	err := database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).First(&cartItem).Error
	if err != nil {
		cartItem = models.CartItem{
			CartID:    cart.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		database.DB.Create(&cartItem)
	} else {
		cartItem.Quantity += input.Quantity
		database.DB.Save(&cartItem)
	}

	return helpers.ResponseSuccess(c, "Product added to cart")
}

func GetCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var cart models.Cart
	if err := database.DB.Preload("CartItems.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		cart = models.Cart{UserID: userID}
		if err := database.DB.Create(&cart).Error; err != nil {
			return helpers.InternalServerError(c, "Failed to create cart")
		}
	}

	var response models.CartDTO
	response.ID = cart.ID
	response.UserID = cart.UserID
	response.CartItems = make([]models.CartItemDTO, len(cart.CartItems))
	for i, item := range cart.CartItems {
		response.CartItems[i] = models.CartItemDTO{
			ID:       item.ID,
			Quantity: item.Quantity,
			CartID:   item.CartID,
			Product: models.ProductDTO{
				ID:          item.Product.ID,
				Name:        item.Product.Name,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				CategoryID:  item.Product.CategoryID,
			},
		}
	}

	return helpers.ResponseSuccessWithData(c, response)
}

func RemoveFromCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input := c.Locals("input").(*models.RemoveFromCartInput)

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return helpers.NotFoundError(c, "Cart not found")
	}

	if err := database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).First(&models.CartItem{}).Error; err != nil {
		return helpers.NotFoundError(c, "Product not found in cart")
	}

	if err := database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).Delete(&models.CartItem{}).Error; err != nil {
		return helpers.InternalServerError(c, "Failed to remove item from cart")
	}

	return helpers.ResponseSuccess(c, "Product removed from cart")
}
