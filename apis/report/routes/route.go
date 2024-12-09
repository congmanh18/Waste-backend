package routes

import (
	handler "smart-waste/apis/report/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupReportRoutes(app *fiber.App, handler handler.ReportHandler) {
	api := app.Group("wastebin/reports")

	api.Post("/", handler.HandlerCreateReport())
	api.Get("/all", handler.HandlerGetAllReports1())
	api.Get("/last/:id", handler.HandlerGetLast())
	api.Get("/allid", handler.HandlerGetAllReports())
	api.Get("/:id", handler.HandlerGetReportByID())
	api.Delete("/:id", handler.HandlerDeleteReport())
	api.Get("/user/:user_id", handler.HandlerGetReportsByUserID())
	api.Get("/wastebin/:wastebin_id", handler.HandlerGetReportsByWasteBinID())
	api.Get("/date", handler.HandlerGetReportsByDate())
}
