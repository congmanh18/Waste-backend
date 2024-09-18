package handler

import (
	"context"

	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerReadWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		wateBinEntity, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, c.Params("id"))
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load wastebin information", false, nil)
			res.SetError(err)
			return res.Send(c)
		}


		res := res.NewRes(fiber.StatusOK, "WasteBin Information: ", true, wateBinEntity)
		return res.Send(c)
	}
}
