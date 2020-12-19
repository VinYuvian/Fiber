package user

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"Go-Fiber/database"
)

func DeleteUser(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id,err := strconv.Atoi(paramId)
	if err!=nil{
		return ctx.Status(fiber.StatusBadRequest).JSON("Can't parse the request")
	}
	for i,user := range database.Users{
		if user.Id == id{
			database.Users = append(database.Users[:i],database.Users[i+1:]...)
		}
	}
	return ctx.Status(fiber.StatusOK).JSON("User removed")
}
