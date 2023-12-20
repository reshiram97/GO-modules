package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/reshiram97/url-shortner/database"
);

func ResolveURL(c *fiber.Ctx) error {

	url:= c.Params("url")
	
	// Create a client with number 0
	r:= database.CreateClient(0)

	// defer its connectionto close at end of stack
	defer r.Close()
	
	// As redis is a key value pair database hence searching for the exact match of any URL served
	value,err := r.Get(database.Ctx,url).Result()

	// err recieved for no match
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No such URL found",
		})
	} else if err!=nil {
		// error recieved for noproper db connection
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "DB not connected",
		})
	}


	// Increament the counter for next subsequent search
	rInr:= database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx,"counter")

	// Redirect to the originla URL
	return c.Redirect(value,301)
}
