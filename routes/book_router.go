// routes.go
package routes

import (
	"github.com/akers1023/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupBooksRoutes(app *fiber.App, repository *handlers.Repository) {
	api := app.Group("/api")
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	api.Post("/create_books", handler.Repository.CreateBook)
	api.Delete("/delete_book/:id", handler.Repository.DeleteBook)
	api.Get("/get_books/:id", handler.Repository.GetBookByID)
	api.Get("/books", handler.Repository.GetBooks)
}
