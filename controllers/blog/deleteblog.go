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

func DeleteBlogPost(c *fiber.Ctx) error {
	blogId := c.Params("id")

	ObjectID, err := primitive.ObjectIDFromHex(blogId)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var blog models.Blog
	err = blogCollection.FindOneAndDelete(context.TODO(), bson.M{"_id": ObjectID}).Decode(&blog)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "The blog id does not exist"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Blog has been successfully delted."})
}
