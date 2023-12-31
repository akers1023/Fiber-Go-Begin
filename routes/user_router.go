package routes

import (
	"github.com/akers1023/handlers"
	"github.com/akers1023/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUsersRoutes(app *fiber.App, repository *handlers.Repository) {
	// api := app.Group("/users")
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	app.Get("/users/:id", middleware.Authenticate(), handler.Repository.GetUser)
	app.Get("/users", middleware.Authenticate(), handler.Repository.GetAllUsers)
}
