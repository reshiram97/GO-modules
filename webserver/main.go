package main

import (
	"github.com/gofiber/fiber/v2"
	"webserver/routes"
	"webserver/database"
);

func routeHandler(app *fiber.App){
	app.Get("/",routes.TestRoute)
	app.Get("/getusers",routes.FindAllUsers)
	app.Get("/getuser/:id",routes.FindUserById)
	app.Post("/createuser",routes.AddUser)
	app.Delete("/removeuser/:id",routes.DeleteUser)
}
func main() {
	app:= fiber.New();
	database.Init()
	routeHandler(app)
	app.Listen(":5000")
}