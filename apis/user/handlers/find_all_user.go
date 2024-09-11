package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) HandlerFindAllUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		userList, err := u.FindAllUserUsecase.ExecuteFindAll(ctx)

		if err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Unable to load user list", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "User List: ", true, userList)
		return res.Send(c)
	}
}
