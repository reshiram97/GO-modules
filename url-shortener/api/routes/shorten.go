package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/reshiram97/url-shortner/database"
	"github.com/reshiram97/url-shortner/helpers"
);

type request struct {
	URL				string			`json:"url"`
	CustomShort		string			`json:"shorten"`
	Expiry			time.Duration	`json:"expiry"`
}

type response struct {
	URL				string			`json:"url"`
	CustomShort		string			`json:"shorten"`
	Expiry			time.Duration	`json:"expiry"`
	XRateRemaining	int				`json:"rate_remaining"`
	XRateLimitReset	time.Duration	`json:"rate_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body:= new(request)
	if err:= c.BodyParser(&body); err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Data cannot be parsed properly"})
	}

	// Implement Rate Limiter

	r2:= database.CreateClient(1)
	defer r2.Close()

	value, err := r2.Get(database.Ctx,c.IP()).Result()

	if err== redis.Nil {
		_ = r2.Set(database.Ctx,c.IP(),os.Getenv("API_QUOTA"),30*60*time.Second).Err()
	} else {
		valInt,_:= strconv.Atoi(value)
		if valInt<=0 {
			limit,_ := r2.TTL(database.Ctx,c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":"Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute
			})
		}
	}

	// Check is its a valid URL
	if !govalidator.isURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL Provided"})
	}

	// Edge Case: If the given URL that is localhost:5000 is not looped in
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"Hacking Attempted"})
	}

	// Enforce HTTP SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	r2.Decr(database.Ctx,c.ID())
}