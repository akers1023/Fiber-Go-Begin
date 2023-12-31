package handlers

import (
	"net/http"
	"time"

	"github.com/akers1023/models"
	"github.com/akers1023/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *Repository) HandlerRegister(context *fiber.Ctx) error {
	user := models.User{}
	user.ID = uuid.New().String()

	err := context.BodyParser(&user)
	if err != nil {
		return HandleErrorResponse(context, http.StatusUnprocessableEntity, "Request failed")
	}

	validationErr := utils.Validate.Struct(user)
	if validationErr != nil {
		return HandleErrorResponse(context, http.StatusBadRequest, "Validation failed")
	}

	//check email already exists?

	token, refreshToken, _ := utils.GenerateAllTokens(*user.Email, user.ID, *user.Role)
	user.Token = &token
	user.RefreshToken = &refreshToken

	password := utils.HashPassword(*user.Password)
	user.Password = &password

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = r.DB.Create(&user).Error
	if err != nil {
		return HandleErrorResponse(context, http.StatusBadRequest, "Could not create user")
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":       "user has been added",
		"token":         token,
		"refresh_token": refreshToken})
}

func (r *Repository) HandlerLogin(context *fiber.Ctx) error {
	var user models.User
	var foundUser models.User

	err := context.BodyParser(&user)
	if err != nil {
		return HandleErrorResponse(context, http.StatusUnprocessableEntity, "Request failed")
	}

	if err := r.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		return HandleErrorResponse(context, http.StatusUnprocessableEntity, err.Error())
	}

	passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
	if passwordIsValid != true {
		return HandleErrorResponse(context, http.StatusInternalServerError, msg)
	}

	if foundUser.Email == nil {
		return HandleErrorResponse(context, http.StatusInternalServerError, "user not found")
	}

	token, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Email, foundUser.ID, *foundUser.Role)
	utils.UpdateAllTokens(r.DB, foundUser.ID, token, refreshToken)

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"full_name":     "Hello " + foundUser.Full_name,
		"message":       "Login successfully",
		"token":         token,
		"refresh_token": refreshToken})
}

func (r *Repository) GetUser(context *fiber.Ctx) (err error) {
	userID := context.Params("id")

	if err := utils.MatchUserTypeToUID(context, userID); err != nil {
		return HandleErrorResponse(context, http.StatusBadRequest, "Request failed")
	}
	var user models.User
	err = r.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return HandleErrorResponse(context, http.StatusInternalServerError, "user not found")
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user is fetched successfully",
		"data":    user,
	})
	return nil
}

func (r *Repository) GetAllUsers(context *fiber.Ctx) error {
	if err := utils.CheckUserType(context, "ADMIN"); err != nil {
		return HandleErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	users := &[]models.User{}

	err := r.DB.Find(users).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "users fetched successfully",
		"data":    users,
	})
	return nil

}
