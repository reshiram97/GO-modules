package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reshiram97/webserver/database"
	"github.com/reshiram97/webserver/middleware"
	"github.com/reshiram97/webserver/routes"
);

func routeHandler(app *fiber.App){
	app.Use("/api/*", middleware.Protectedhandler)
	app.Get("/api/test",routes.TestRoute)
	app.Get("/getusers",routes.FindAllUsers)
	app.Get("/getuser/:id",routes.FindUserById)
	app.Post("/createuser",routes.RegisterUser)
	app.Post("/login",routes.LoginUser)
	app.Delete("/removeuser/:id",routes.DeleteUser)
	app.Post("/api/createpost",routes.AddPost);
	// app.Get("/posts",routes.FindPosts)
}
func main() {
	app:= fiber.New();
	database.Init()
	routeHandler(app)
	app.Listen(":5000")
}