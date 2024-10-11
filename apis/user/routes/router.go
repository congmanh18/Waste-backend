package routes

import (
	handler "smart-waste/apis/user/handlers"
	"smart-waste/pkgs/middleware"

	_ "smart-waste/docs"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes thiết lập các route cho người dùng
// @title Smart Waste Management API
// @version 1.0
// @description Đây là API cho quản lý người dùng trong hệ thống quản lý chất thải thông minh.
// @host localhost:3000
// @BasePath /users

func SetupUserRoutes(app *fiber.App, userHander handler.UserHandler) {
	var userRoutes = app.Group("/users")

	// Public routes
	// @Summary Đăng nhập
	// @Description Đăng nhập cho người dùng
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param credentials body handler.LoginRequest true "Thông tin đăng nhập"
	// @Success 200 {object} handler.LoginResponse
	// @Failure 400 {object} fiber.Map
	// @Router /users/login [post]
	userRoutes.Post("/login", userHander.HandlerLogin())

	// @Summary Làm mới token
	// @Description Làm mới token đã hết hạn
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Success 200 {object} handler.TokenResponse
	// @Failure 400 {object} fiber.Map
	// @Router /users/refresh [post]
	userRoutes.Post("/refresh", userHander.RefreshTokenHandler())

	// Protected routes
	authRoutes := userRoutes.Group("/", middleware.AuthMiddleware)

	// @Summary Đăng ký người dùng mới
	// @Description Tạo một người dùng mới
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param user body handler.CreateUserRequest true "Thông tin người dùng"
	// @Success 201 {object} handler.UserResponse
	// @Failure 400 {object} fiber.Map
	// @Router /users/register [post]
	authRoutes.Post("/register", userHander.HandlerCreateUser())

	// @Summary Cập nhật thông tin người dùng
	// @Description Cập nhật thông tin của một người dùng đã có
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param id path string true "ID người dùng"
	// @Param user body handler.UpdateUserRequest true "Thông tin cập nhật"
	// @Success 200 {object} handler.UserResponse
	// @Failure 400 {object} fiber.Map
	// @Router /users/{id} [put]
	authRoutes.Put("/:id", userHander.HandlerUpdateUser())

	// @Summary Lấy thông tin người dùng theo ID
	// @Description Lấy thông tin chi tiết của một người dùng theo ID
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param id path string true "ID người dùng"
	// @Success 200 {object} handler.UserResponse
	// @Failure 400 {object} fiber.Map
	// @Router /users/{id} [get]
	authRoutes.Get("/:id", userHander.HandlerFindUserByID())

	// Admin routes
	adminRoutes := app.Group("/admin", middleware.AuthMiddleware, userHander.AdminOnlyHandler())

	// @Summary Lấy danh sách người dùng
	// @Description Chỉ dành cho admin, lấy danh sách toàn bộ người dùng
	// @Tags Admin
	// @Accept json
	// @Produce json
	// @Success 200 {array} handler.UserResponse
	// @Failure 400 {object} fiber.Map
	// @Router /admin/findall [get]
	adminRoutes.Get("/findall", userHander.HandlerFindAllUser())

	// @Summary Xóa người dùng
	// @Description Chỉ dành cho admin, xóa một người dùng theo ID
	// @Tags Admin
	// @Accept json
	// @Produce json
	// @Param id path string true "ID người dùng"
	// @Success 204 {string} string "Xóa thành công"
	// @Failure 400 {object} fiber.Map
	// @Router /admin/{id} [delete]
	adminRoutes.Delete("/:id", userHander.HandlerDeleteUser())
}
