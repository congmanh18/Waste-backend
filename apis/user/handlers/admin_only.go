package handler

import (
	"smart-waste/pkgs/res"

	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) AdminOnlyHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)

		if role != "admin" {
			res := res.NewRes(fiber.StatusForbidden, "Access denied", false, nil)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "Welcome, admin!", true, nil)
		return res.Send(c)
	}
}
