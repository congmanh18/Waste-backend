package routes

import (
	handler "smart-waste/apis/wastebin/handlers"

	_ "smart-waste/docs"

	"github.com/gofiber/fiber/v2"
)

// SetupWasteBinRoutes thiết lập các route cho wastebin
// @title Smart Waste Management API - Waste Bin
// @version 1.0
// @description API để quản lý wastebin trong hệ thống quản lý chất thải thông minh.
// @host localhost:3000
// @BasePath /wastebin

func SetupWasteBinRoutes(app *fiber.App, wasteBinHandler handler.WasteBinHandler) {
	var binRoutes = app.Group("/wastebin")

	// @Summary Kết nối WebSocket
	// @Description Kết nối tới wastebin qua WebSocket
	// @Tags WasteBin
	// @Produce json
	// @Success 101 {string} string "WebSocket connected"
	// @Failure 400 {object} fiber.Map
	// @Router /wastebin/ws [get]
	binRoutes.Get("/ws", wasteBinHandler.WebSocketHandler())

	// @Summary Tạo wastebin mới
	// @Description Tạo một wastebin mới với thông tin cụ thể
	// @Tags WasteBin
	// @Accept json
	// @Produce json
	// @Param wastebin body handler.CreateWasteBinRequest true "Thông tin wastebin"
	// @Success 201 {object} handler.WasteBinResponse
	// @Failure 400 {object} fiber.Map
	// @Router /wastebin [post]
	binRoutes.Post("/", wasteBinHandler.HandlerCreateWasteBin())

	// @Summary Cập nhật thông tin wastebin
	// @Description Cập nhật thông tin của một wastebin theo ID
	// @Tags WasteBin
	// @Accept json
	// @Produce json
	// @Param id path string true "ID wastebin"
	// @Param wastebin body handler.UpdateWasteBinRequest true "Thông tin cập nhật"
	// @Success 200 {object} handler.WasteBinResponse
	// @Failure 400 {object} fiber.Map
	// @Router /wastebin/{id} [put]
	binRoutes.Put("/:id", wasteBinHandler.HandlerUpdateWasteBin())

	// @Summary Xóa wastebin
	// @Description Xóa một wastebin theo ID
	// @Tags WasteBin
	// @Produce json
	// @Param id path string true "ID wastebin"
	// @Success 204 {string} string "Xóa thành công"
	// @Failure 400 {object} fiber.Map
	// @Router /wastebin/{id} [delete]
	binRoutes.Delete("/:id", wasteBinHandler.HandlerDeleteWasteBin())

	// @Summary Lấy thông tin wastebin
	// @Description Lấy thông tin của một wastebin theo ID
	// @Tags WasteBin
	// @Produce json
	// @Param id path string true "ID wastebin"
	// @Success 200 {object} handler.WasteBinResponse
	// @Failure 400 {object} fiber.Map
	// @Router /wastebin/{id} [get]
	binRoutes.Get("/:id", wasteBinHandler.HandlerReadWasteBin())
}
