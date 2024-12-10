package handlers

import (
	"errors"

	"github.com/irfanalmsyah/ecommerce-api/database"
	"github.com/irfanalmsyah/ecommerce-api/helpers"
	"github.com/irfanalmsyah/ecommerce-api/models"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	input := c.Locals("input").(*models.RegisterInput)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return helpers.InternalServerError(c, "Failed to hash password")
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return helpers.ConflictError(c, "Email already exists")
		}
		return helpers.InternalServerError(c, "Failed to create user")
	}

	return helpers.ResponseCreated(c, "User registered successfully")
}

func Login(c *fiber.Ctx) error {
	input := c.Locals("input").(*models.LoginInput)

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return helpers.UnauthorizedError(c, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return helpers.UnauthorizedError(c, "Invalid credentials")
	}

	tokenString, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return helpers.InternalServerError(c, "Failed to generate token")
	}

	return helpers.ResponseSuccessWithData(c, fiber.Map{"token": tokenString})
}
