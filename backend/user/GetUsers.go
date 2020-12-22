package user

import (
	"Go-Fiber/database"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error{
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Users":database.Users,
		//"jwt" : ctx.Locals("user").(*jwt.Token),
	})
}
