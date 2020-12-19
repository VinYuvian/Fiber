package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"Go-Fiber/user"
)


func main() {
	app := fiber.New()
	file,err := os.OpenFile("log.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalf("error opening file : %v",err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output : file,
		Format : "[${time}] ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HELLO")
	})
	app.Post("/User/upload",user.UserUpload)
	app.Get("/Users",user.GetUsers)
	app.Post("/User",user.CreateUser)
	app.Get("/Users/:id",user.GetUser)
	app.Delete("/Users/:id",user.DeleteUser)
	err = app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}


