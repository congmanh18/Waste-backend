package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerReadAllIDWasteBins() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		wasteBinEntities, err := w.ReadAllWasteBinUsecase.ReadAllWasteBins(ctx)
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load wastebin information", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		var wasteBinIDs []string
		for _, report := range wasteBinEntities {
			wasteBinIDs = append(wasteBinIDs, report.ID)
		}

		res := res.NewRes(fiber.StatusOK, "All WasteBin Information", true, wasteBinIDs)
		return res.Send(c)
	}
}
