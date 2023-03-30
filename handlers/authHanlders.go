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
	Username  string `json:"username"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Instance  string `json:"instance"`
	Workgroup string `json:"workgroup"`
	Address   struct {
		Province string `json:"province"`
		City     string `json:"city"`
		District string `json:"district"`
		Ward     string `json:"ward"`
	} `json:"address"`
}

type loginStruct struct {
	Username string
	Password string
}

type responseStruct struct {
	Rescode int    `json:"rescode"`
	Message string `json:"message"`
	Data    struct {
		User  userStruct `json:"user"`
		Token string     `json:"token"`
	} `json:"data"`
}

func Signin(c *fiber.Ctx) error {
	u := new(loginStruct)

	if err := c.BodyParser(u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"rescode": 400,
			"success": false,
			"message": err.Error(),
		})
	}

	var user models.User
	initializers.DB.Preload("UserRole").Preload("UserInstance").Preload("WorkGroup").Preload("Province").Preload("City").Preload("District").Preload("Ward").Find(&user, "username = ?", u.Username)

	if user.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"rescode": 400,
			"success": false,
			"message": "username and password not match",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"rescode": 400,
			"success": false,
			"message": "username and password not match",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"rol": user.UserRoleId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"rescode": 400,
			"success": false,
			"message": err.Error(),
		})
	}

	responseOK := responseStruct{
		Rescode: 200,
		Message: "success get data",
		Data: struct {
			User  userStruct "json:\"user\""
			Token string     "json:\"token\""
		}{
			User:  formatUser(user),
			Token: tokenString},
	}
	return c.Status(200).JSON(responseOK)
}

func Validate(c *fiber.Ctx) error {

	sub := c.Locals("sub")

	var user models.User
	initializers.DB.Preload("UserRole").Preload("UserInstance").Preload("WorkGroup").Preload("Province").Preload("City").Preload("District").Preload("Ward").Find(&user, sub)

	return c.Status(200).JSON(&fiber.Map{
		"rescode": 200,
		"success": true,
		"user":    formatUser(user),
	})
}

func formatUser(user models.User) userStruct {
	newFormat := userStruct{
		Username:  user.Username,
		Name:      user.Name,
		Role:      user.UserRole.RoleName,
		Instance:  user.UserInstance.InstanceName,
		Workgroup: user.WorkGroup.GroupName,
		Address: struct {
			Province string "json:\"province\""
			City     string "json:\"city\""
			District string "json:\"district\""
			Ward     string "json:\"ward\""
		}{
			Province: user.Province.Name,
			City:     user.City.Name,
			District: user.District.Name,
			Ward:     user.Ward.Name,
		},
	}

	return newFormat
}
