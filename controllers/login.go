package controllers

import (
	"app/models"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func LoginUser(c *fiber.Ctx) error {
	user := struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	var foundUser models.User
	filter := bson.M{"email": user.Email}
	err := userCollection.FindOne(context.Background(), filter).Decode(&foundUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "No user found with the given email"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid Email"})
	}

	if !CheckPassword(user.Password, foundUser.Password) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "The password is incorrect for the given email."})
	}

	claims := jwt.MapClaims{
		"username": foundUser.Username,
		"email":    foundUser.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate a token. Please contact the dev."})
	}

	return c.JSON(fiber.Map{"user": user, "token": tokenString})
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
