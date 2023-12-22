package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reshiram97/webserver/routes"
	"github.com/reshiram97/webserver/database"
	"github.com/google/uuid"
);

func routeHandler(app *fiber.App){
	app.Use("/api/*", func (c *fiber.Ctx) error {
		userUUID:= c.Get("Authorization")
		if userUUID==""{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Access"})
		}
		reuserUUID,err := uuid.Parse(userUUID)
		if err!=nil {
			return c.Status(500).SendString("Wrong Parse")
		}
		c.Locals("userUUID", reuserUUID)
		// Proceed to the next middleware or route handler
		return c.Next()
	})
	app.Get("/",routes.TestRoute)
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