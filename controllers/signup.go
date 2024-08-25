package controllers

import (
	"context"
	"net/http"
	"time"

	"app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func SetUserController(c *mongo.Collection) {
	userCollection = c
}

func SignUp(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var emailExists models.User
	filter := bson.M{"email": user.Email}
	findUserErr := userCollection.FindOne(context.Background(), filter).Decode(&emailExists)

	if findUserErr != nil && findUserErr != mongo.ErrNoDocuments {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": findUserErr.Error()})
	}

	if findUserErr == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "This email already exists"})
	}

	hashedPassword, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": hashErr.Error()})
	}
	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, insertErr := userCollection.InsertOne(ctx, user)

	if insertErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": insertErr.Error()})
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could note generate token."})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"user": user, "token": tokenString})
}

func UserDoesNotExist(ctx context.Context, email string) bool {
	var user models.User
	filter := bson.M{"email": user.Email}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	return err == nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
