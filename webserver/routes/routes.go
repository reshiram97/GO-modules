package routes

import (
	"fmt"
	"webserver/database"
	"webserver/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
);
var db = database.Init()

func createNewUUID() uuid.UUID {
	return uuid.New()
}

func TestRoute(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func AddUser(c *fiber.Ctx) error {
	var user models.User
	if err:= c.BodyParser(&user); err!=nil {
		return err
	}
	user.Id = createNewUUID()
	db.Create(&user)
	return c.Status(200).SendString("User Created Successfully") 
}

func DeleteUser(c *fiber.Ctx) error {
	id:= c.Params("id")
	fmt.Println(id)
	var user models.User
	if err:= db.Find(&user,"id=?",id).Error; err!=nil {
		return c.Status(404).SendString("User not Found")
	}
	if err:= db.Unscoped().Delete(&user).Error; err!=nil {
		return c.Status(500).SendString("User cannot be removed")
	}
	return c.Status(200).SendString("User Deleted Successfully")
}

func FindAllUsers(c *fiber.Ctx) error {	
	var user []models.User
	db.Find(&user)
	return c.Status(200).JSON(user) 
}

func FindUserById(c *fiber.Ctx) error {
	id:= c.Params("id")
	var user models.User
	if err:= db.First(&user,"id=?",id).Error; err!=nil {
		return c.Status(404).SendString("User not found")
	}
	return c.Status(200).JSON(user)
}