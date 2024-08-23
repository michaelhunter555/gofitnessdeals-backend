package controllers

import (
	"context"
	"net/http"

	"app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	ObjectID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Could not find user by id"})
	}

	var user models.User
	err = userCollection.FindOneAndDelete(context.TODO(), bson.M{"_id": ObjectID}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found in database"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "There was an internal server error"})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User deleted successfully."})
}
