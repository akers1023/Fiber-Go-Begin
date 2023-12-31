package routes

import (
	"github.com/akers1023/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, repository *handlers.Repository) {
	api := app.Group("/users")
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	api.Post("/signup", handler.Repository.HandlerRegister)
	// api.Post("/signin", middleware.Authenticate(), handler.Repository.HandlerLogin)
	api.Post("/signin", handler.Repository.HandlerLogin)
}
