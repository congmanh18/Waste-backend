package routes

import (
	handler "smart-waste/apis/wastebin/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupWasteBinRoutes(app *fiber.App, wasteBinHandler handler.WasteBinHandler) {
	var binRoutes = app.Group("/wastebins")

	binRoutes.Post("/", wasteBinHandler.HandlerCreateWasteBin())
	binRoutes.Put("/:id", wasteBinHandler.HandlerUpdateWasteBin())
	binRoutes.Delete("/:id/remove", wasteBinHandler.HandlerDeleteWasteBin())
	binRoutes.Get("/:id/info", wasteBinHandler.HandlerReadWasteBin())
	binRoutes.Get("/", wasteBinHandler.HandlerReadAllIDWasteBins())
	binRoutes.Get("/all", wasteBinHandler.HandlerReadAllWasteBins())
}

func SetupExponentialSmoothingRoutes(app *fiber.App, wasteBinHandler handler.WasteBinHandler) {
	app.Get("/models/expo/:id", wasteBinHandler.HandlerExponentialSmoothing())
	app.Get("/models/svm/:id", wasteBinHandler.HandlerSVM())
}
