package routes

import (
	handler "smart-waste/apis/user/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHander handler.UserHandler) {
	var userRoutes = app.Group("/wastebin/users")
	userRoutes.Post("/login", userHander.HandlerLogin())          // Create User
	userRoutes.Post("/refresh", userHander.RefreshTokenHandler()) // Create User
	userRoutes.Post("/register", userHander.HandlerCreateUser())  // Create User
	userRoutes.Put("/:id", userHander.HandlerUpdateUser())        // Create User
	userRoutes.Get("/:id", userHander.HandlerFindUserByID())      // Create User
	userRoutes.Get("/", userHander.HandlerFindAllUser())          // Create User
	userRoutes.Delete("/:id", userHander.HandlerDeleteUser())     // Create User
}
