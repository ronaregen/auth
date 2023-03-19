package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ronaregen/auth/initializers"
	"github.com/ronaregen/auth/models"
	"golang.org/x/crypto/bcrypt"
)

type userStruct struct {
	Username string
	Name     string
	Password string
}

type loginStruct struct {
	Username string
	Password string
}

func Signup(c *fiber.Ctx) error {
	u := new(userStruct)

	if err := c.BodyParser(u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	user := models.User{Username: u.Username, Name: u.Name, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": result.Error.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"user":    user,
	})
}

func Signin(c *fiber.Ctx) error {
	u := new(loginStruct)

	if err := c.BodyParser(u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	var user models.User
	initializers.DB.Find(&user, "username = ?", u.Username)

	if user.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "username and password not match",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "username and password not match",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"token":   tokenString,
	})
}

func Validate(c *fiber.Ctx) error {

	user := c.Locals("user")

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"user":    user,
	})
}
