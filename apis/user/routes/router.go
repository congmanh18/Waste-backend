package routes

import (
	handler "smart-waste/apis/user/handlers"
	"smart-waste/pkgs/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHander handler.UserHandler) {
	var userRoutes = app.Group("/users")

	// Public routes
	userRoutes.Post("/login", userHander.HandlerLogin())

	// Protected routes
	authRoutes := userRoutes.Group("/", auth.AuthMiddleware)
	authRoutes.Post("/register", userHander.HandlerCreateUser())
	authRoutes.Put("/:id", userHander.HandlerUpdateUser())
	authRoutes.Get("/:id", userHander.HandlerFindUserByID())

	// Admin routes (add /admin to differentiate)
	adminRoutes := app.Group("/admin", auth.AuthMiddleware, userHander.AdminOnlyHandler())
	adminRoutes.Get("/findall", userHander.HandlerFindAllUser())
	adminRoutes.Delete("/:id", userHander.HandlerDeleteUser())

	// adminRoutes.Post("/refresh", userHander.RefreshTokenHandler()) ??????
}
