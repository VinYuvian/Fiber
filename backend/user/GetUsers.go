package user

import (
	"github.com/gofiber/fiber/v2"
	"Go-Fiber/database"
)

func GetUsers(ctx *fiber.Ctx) error{
	return ctx.Status(fiber.StatusOK).JSON(database.Users)
}
