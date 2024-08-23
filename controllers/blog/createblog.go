package blog

import (
	"context"
	"net/http"
	"time"

	"app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var blogCollection *mongo.Collection

func SetBlogCollection(c *mongo.Collection) {
	blogCollection = c
}

func CreateBlogPost(c *fiber.Ctx) error {
	var blog models.Blog

	if err := c.BodyParser(&blog); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	blog.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, addBlogErr := blogCollection.InsertOne(ctx, blog)

	if addBlogErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": addBlogErr.Error()})
	}

	return c.Status(http.StatusCreated).JSON(blog)
}
