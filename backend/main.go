package main

import (
	"fmt"
	"github.com/VinYuvian/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)


func main() {
	err:=godotenv.Load(".env")
	if err!=nil{
		fmt.Println("error loading .env file")
	}
	app := fiber.New()
	file,err := os.OpenFile("log.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalf("error opening file : %v",err)
	}
	defer file.Close()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Output : file,
		Format : "[${time}] ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HELLO")
	})
	app.Post("/User/upload", handlers.UserUpload)
	//app.Get("/Users",handlers.AuthRequired(),handlers.GetUsers)
	app.Get("/Users", handlers.GetUsers)
	app.Post("/Signup", handlers.CreateUser)
	app.Get("/Users/:email", handlers.GetUser)
	app.Delete("/Users/:email", handlers.DeleteUser)
	app.Post("/Login", handlers.Login)
	err = app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}


