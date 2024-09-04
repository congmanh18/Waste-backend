package routes

import (
	handler "smart-waste/apis/user/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHander handler.UserHandler) {
	// Tạo một instance của UserHandler, đang để ở main
	// Định nghĩa route cho việc tạo user
	app.Post("/users", userHander.HandlerCreateUser())
	app.Put("/users/:id", userHander.HandlerUpdateUser())
	app.Delete("/users/:id", userHander.HandlerDeleteUser())
	app.Get("/users/:id", userHander.HandlerFindUserByID())
	app.Get("/users", userHander.HandlerFindAllUser())
}
