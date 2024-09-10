package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) HandlerFindUserByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		userEntity, err := u.FindUserByIDUsecase.ExecuteFindById(ctx, c.Params("id"))
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load user information", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "User Information: ", true, userEntity)
		return res.Send(c)
	}
}
