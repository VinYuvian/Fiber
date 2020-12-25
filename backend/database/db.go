package database
import(
	"Go-Fiber/models"
)

var Users = []models.User{
	{
		Id:        1,
		FirstName: "Vinay",
		LastName:  "Nuthipelly",
		Gender:    "Male",
		Email:     "vinay.nuthipelly@gmail.com",
		Password:  "Vin@y1711",
	},
	{
		Id:        2,
		FirstName: "Dileep",
		LastName:  "Nuthipelly",
		Gender:    "Male",
		Email:     "dileep.nuthipelly@gmail.com",
		Password:  "Vin@y1711",
	},
}
