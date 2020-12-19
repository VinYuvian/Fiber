package user

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"Go-Fiber/database"
)

func GetUser(ctx *fiber.Ctx) error{
	paramsId := ctx.Params("id")
	id,err :=strconv.Atoi(paramsId)
	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Can't parse the request",
		})
	}
	for _,user := range database.Users{
		if user.Id == id{
			ctx.Status(fiber.StatusOK).JSON(user)
			return nil
		}
	}
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message" : "data not found",
	})

}
