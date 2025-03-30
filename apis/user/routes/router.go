package routes

import (
	handler "smart-waste/apis/user/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHander handler.UserHandler) {
	var userRoutes = app.Group("/users")
	userRoutes.Post("/login", userHander.HandlerLogin())             // Create User
	userRoutes.Post("/refresh", userHander.RefreshTokenHandler())    // Create User
	userRoutes.Post("/register", userHander.HandlerCreateUser())     // Create User
	userRoutes.Get("/all", userHander.HandlerFindAllUser())          // Create User
	userRoutes.Delete("/:id/remove", userHander.HandlerDeleteUser()) // Create User
	userRoutes.Put("/:id/edit", userHander.HandlerUpdateUser())      // Create User
	userRoutes.Get("/:id/info", userHander.HandlerFindUserByID())
}
