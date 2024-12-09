package routes

import (
	handler "smart-waste/apis/wastebin/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupWasteBinRoutes(app *fiber.App, wasteBinHandler handler.WasteBinHandler) {
	var binRoutes = app.Group("/wastebin")

	binRoutes.Post("/", wasteBinHandler.HandlerCreateWasteBin())
	binRoutes.Put("/:id", wasteBinHandler.HandlerUpdateWasteBin())
	binRoutes.Delete("/:id", wasteBinHandler.HandlerDeleteWasteBin())
	binRoutes.Get("/:id", wasteBinHandler.HandlerReadWasteBin())
	binRoutes.Get("/", wasteBinHandler.HandlerReadAllIDWasteBins())
	binRoutes.Get("/info/all", wasteBinHandler.HandlerReadAllWasteBins())
	binRoutes.Get("/ws/update", websocket.New(wasteBinHandler.WebSocketUpdateWasteBin))

}

func SetupExponentialSmoothingRoutes(app *fiber.App, wasteBinHandler handler.WasteBinHandler) {
	app.Get("/wastebin/ml/exponential", handler.HandlerExponentialSmoothing())
	app.Get("/wastebin/ml/svm/:id", wasteBinHandler.HandlerSVM())
}
