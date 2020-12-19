package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"Go-Fiber/types"
	"Go-Fiber/database"
)

func CreateUser(ctx *fiber.Ctx) error{
	file,err := ctx.FormFile("document")
	if err!=nil{
		return err
	}
	var body types.User
	err = ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "can't parse json",
		})
	}
	user := types.User{
		Id:        len(database.Users)+1,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Gender:    body.Gender,
		Email:     body.Email,
		Password:  body.Password,
	}
	database.Users = append(database.Users,user)
	ctx.SaveFile(file,fmt.Sprintf("./uploads/%s",strconv.Itoa(user.Id)+file.Filename))
	return ctx.Status(fiber.StatusCreated).JSON(user)
}
