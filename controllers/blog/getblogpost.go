package blog

import (
	"context"
	"net/http"

	"app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBlogPost(c *fiber.Ctx) error {
	blogId := c.Params("blogId")

	ObjectID, err := primitive.ObjectIDFromHex(blogId)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var blog models.Blog
	findBlogErr := blogCollection.FindOne(context.TODO(), bson.M{"_id": ObjectID}).Decode(&blog)

	if findBlogErr == nil {
		if findBlogErr == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No user found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(http.StatusCreated).JSON(blog)
}
