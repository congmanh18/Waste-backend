package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerDeleteWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		err := w.DeleteWasteBinUsecase.ExecuteDeleteWasteBin(ctx, c.Params("id"))
		if err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to delete trastbin", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "Waste Bin deleted successfully", true, nil)
		return res.Send(c)
	}
}
