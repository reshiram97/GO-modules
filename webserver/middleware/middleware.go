package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reshiram97/webserver/jwt"
)

func Protectedhandler(c *fiber.Ctx) error {
	tokenString:= c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Not Logged In"})
	}
	tokenString = tokenString[len("Bearer "):]
	if err:= jwt.VerifyToken(tokenString); err!=nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid Token"})
	}
	return c.Next()
} 