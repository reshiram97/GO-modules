package routes
import (
	"fmt"
	"github.com/reshiram97/webserver/database"
	"github.com/reshiram97/webserver/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/reshiram97/webserver/jwt"
);
var db = database.Init()

type loginRequestData struct {
	Email	string 	`json:"email"`
	Password	string  `json:"password"`
}

func createNewUUID() uuid.UUID {
	return uuid.New()
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func TestRoute(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err:= c.BodyParser(&user); err!=nil {
		return err
	}
	user.Id = createNewUUID()
	user.Password,_ = HashPassword(user.Password)
	db.Create(&user)
	return c.Status(200).SendString("User Created Successfully") 
}

func LoginUser(c *fiber.Ctx) error {
	var login loginRequestData
	var user models.User
	if err:= c.BodyParser(&login); err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err})
	}
	if err:= db.Table("users").Where("email=?",login.Email).First(&user).Error; err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err})
	}
	if err:= CheckPasswordHash(login.Password,user.Password); err==false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err})
	}
	tokenString,err := jwt.CreateToken(user.Email)
	if err!=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"No matching email found"})
	}
	c.Set("Authorization",tokenString)
	return c.Status(200).SendString(tokenString)
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

func AddPost(c *fiber.Ctx) error {
	var user models.User
	email,_ := jwt.ExtractClaims(c.Get("Authorization"))
	if email=="Invalid" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized - User not logged in")
	}
	var post models.Post
	if err:= c.BodyParser(&post); err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err})
	}
	if err:= db.First(&user,"email=?",email).Error; err!=nil {
		return c.Status(404).SendString("User not found")
	}
	post.User = user.Id
	post.Id = createNewUUID()
	db.Create(&post)
	return c.Status(200).JSON(fiber.Map{"Status":"Post Created Success"})
}

