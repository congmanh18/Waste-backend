package routes

import (
	"smart-waste/apis/wastebin/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupWasteBinRoutes(app *fiber.App, wasteBinHandler handler.WasteBinHandler) {
	app.Post("/bin", wasteBinHandler.HandlerCreateWasteBin())
	app.Put("/bin/:id", wasteBinHandler.HandlerUpdateWasteBin())
	app.Delete("/bin/:id", wasteBinHandler.HandlerDeleteWasteBin())
	app.Get("/bin/:id", wasteBinHandler.HandlerReadWasteBin())
}
