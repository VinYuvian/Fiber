package user

import (
	"Go-Fiber/database"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"time"
)
const secretKey ="userKey1711"
func AuthRequired() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey: []byte(secretKey),
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}

func Login(ctx *fiber.Ctx) error{
	type cred struct{
		Email string
		Password string
	}
	var credentials cred
	err := ctx.BodyParser(&credentials)
	if err!=nil{
		ctx.Send([]byte("Credentials are required"))
	}
	for _,v:= range database.Users{
		if credentials.Email == v.Email && credentials.Password == v.Password{
			token:=jwt.New(jwt.SigningMethodHS256) //creating new token
			claims:=token.Claims.(jwt.MapClaims) //adding claims
			claims["user"] = v // assigning claims
			claims["exp"] = time.Now().Add(time.Hour*1)
			t,err:=token.SignedString([]byte(secretKey)) //signing the token with the key
			if err!=nil{
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}
			ctx.Status(fiber.StatusOK).JSON(fiber.Map{ //sending the token as response
				"token":t,
			})
			break
		}
	}
	return nil
}
