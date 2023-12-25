package handler

import (
	"fiber-begin/models"
	"fiber-begin/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var validate = validator.New()

func (ur *UserRepository) HandleRegister(context *fiber.Ctx) error {
	// Parse request body
	user := models.User{}
	if err := context.BodyParser(&user); err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Failed to parse the request body",
		})
	}

	// Validate user input
	if err := validate.Struct(user); err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Validation failed",
			"errors":  err.(validator.ValidationErrors),
		})
	}

	// Check if email already exists
	if err := ur.DB.Where("email = ?", user.Email).First(&models.User{}).Error; err == nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Email already exists",
		})
	}

	// Hash the password securely
	password := utils.HashPassword(*user.Password)
	user.Password = &password

	user.Created_at, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
	user.Updated_at, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))

	token, refreshToken, _ := utils.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	// Create user in the database
	if err := ur.DB.Create(&user).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Failed to create user",
		})
	}

	return context.JSON(&fiber.Map{
		"message": "Registration successful",
		"token":   user.Token,
	})
}

func (u *UserRepository) HanldleLogin(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(utils.Resp(true, "sign in!", nil, nil))
}

// func Profile(c *fiber.Ctx) error {
// 	return c.Status(fiber.StatusOK).JSON(utils.Resp(true, "sign in!", nil, nil))
// }
// func UpdateProfile(c *fiber.Ctx) error {
// 	return c.Status(fiber.StatusOK).JSON(utils.Resp(true, "sign in!", nil, nil))
// }
