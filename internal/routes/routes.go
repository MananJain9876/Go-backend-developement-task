package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/example/user-age-api/internal/handler"
	"github.com/example/user-age-api/internal/service"
)

func RegisterUserRoutes(app *fiber.App, userService service.UserService, logger *zap.Logger) {
	h := handler.NewUserHandler(userService, logger)

	app.Post("/users", h.CreateUser)
	app.Get("/users/:id", h.GetUser)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
	app.Get("/users", h.ListUsers)
}


