package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reshiram97/webserver/database"
	"github.com/reshiram97/webserver/middleware"
	"github.com/reshiram97/webserver/routes"
);

func routeHandler(app *fiber.App){
	// Middlewares
	app.Use("/api/*", middleware.Protectedhandler)

	// Admin Routes
	app.Get("/",routes.TestRoute)
	app.Get("/getusers",routes.FindAllUsers)
	app.Get("/getuser/:id",routes.FindUserById)
	app.Delete("/removeuser/:id",routes.DeleteUser)
	
	// User Routes
	app.Post("/register",routes.RegisterUser)
	app.Post("/login",routes.LoginUser)
	app.Get("/logout",routes.LogoutUser)
	
	app.Post("/api/create",routes.AddPost)
	app.Delete("/api/remove/:id",routes.RemovePost)
	app.Put("/api/update",routes.UpdatePost)
	app.Get("/api/posts",routes.GetPosts)
	app.Get("/api/post/:id",routes.GetPost)
}
func main() {
	app:= fiber.New();
	database.Init()
	routeHandler(app)
	app.Listen(":5000")
}