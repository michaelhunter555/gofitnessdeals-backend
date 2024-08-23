package controllers

import (
	"context"
	"net/http"

	"app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
)

var newUserCollection *mongo.Collection

func GetUser(c *fiber.Ctx) error {
	//req.body password, name, email, etc.
	userId := c.Params("id")

	ObjectID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User Id"})
	}

	var user models.User
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": ObjectID}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No user found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to find user"})
	}

	return c.Status(http.StatusCreated).JSON(user)
}
