package routes

import (
	handler "smart-waste/apis/user/handlers"
	"smart-waste/pkgs/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHander handler.UserHandler) {
	var userRoutes = app.Group("/users")

	// Public routes
	userRoutes.Post("/login", userHander.HandlerLogin())
	userRoutes.Post("/refresh", userHander.RefreshTokenHandler())

	// Protected routes
	authRoutes := userRoutes.Group("/", middleware.AuthMiddleware)
	authRoutes.Post("/register", userHander.HandlerCreateUser())
	authRoutes.Put("/:id", userHander.HandlerUpdateUser())
	authRoutes.Get("/:id", userHander.HandlerFindUserByID())

	// Admin routes (add /admin to differentiate)
	adminRoutes := app.Group("/admin", middleware.AuthMiddleware, userHander.AdminOnlyHandler())
	adminRoutes.Get("/findall", userHander.HandlerFindAllUser())
	adminRoutes.Delete("/:id", userHander.HandlerDeleteUser())

}
