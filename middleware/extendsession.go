package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_TOKEN"))

func ExtendSession(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	now := time.Now().Unix()
	exp := int64(claims["exp"].(float64))

	if exp-now < 600 {
		newExp := time.Now().Add(30 * time.Minute).Unix()
		claims["exp"] = newExp

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, tokenErr := newToken.SignedString(jwtSecret)

		if tokenErr != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": tokenErr.Error()})
		}

		c.Set("Authorization", "Bearer "+tokenString)
	}

	return c.Next()
}
