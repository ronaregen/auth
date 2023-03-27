package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ronaregen/auth/initializers"
	"github.com/ronaregen/auth/models"
)

type requestHeader struct {
	Authorization string
}

func ReqAuth(c *fiber.Ctx) error {
	r := new(requestHeader)
	if err := c.ReqHeaderParser(r); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "unauthorized",
		})
	}
	tokenAuth := strings.Fields(r.Authorization)
	if tokenAuth[0] != "Bearer" {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "unauthorized",
		})
	}

	tokenString := tokenAuth[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"success": false,
				"message": "unauthorized",
			})
		}

		var user models.User
		initializers.DB.Find(&user, claims["sub"])
		if user.ID == 0 {
			return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"success": false,
				"message": "unauthorized",
			})
		}

		c.Locals("sub", claims["sub"])
		return c.Next()
	} else {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "unauthorized with error " + err.Error(),
		})
	}

}
