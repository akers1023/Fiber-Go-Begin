package handlers

import (
	"fmt"
	"net/http"

	"github.com/akers1023/models"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Handler struct {
	Repository *Repository
}

func HandleErrorResponse(context *fiber.Ctx, statusCode int, message string) error {
	return context.Status(statusCode).JSON(&fiber.Map{"message": message})
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := models.Books{}
	book.ID = uuid.New().String()
	err := context.BodyParser(&book)
	if err != nil {
		return HandleErrorResponse(context, http.StatusUnprocessableEntity, "Request failed")
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		return HandleErrorResponse(context, http.StatusBadRequest, "Could not create book")
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book has been added"})
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).Delete(&bookModel)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book delete successfully",
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &models.Books{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book id fetched successfully",
		"data":    bookModel,
	})
	return nil
}
