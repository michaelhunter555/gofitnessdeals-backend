package blog

import (
	"context"
	"net/http"

	"app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBlogPosts(c *fiber.Ctx) error {
	var blogPosts []*models.Blog

	curr, err := blogCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer curr.Close(context.TODO())

	for curr.Next(context.TODO()) {
		var blogPost models.Blog

		if err := curr.Decode(&blogPost); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		blogPosts = append(blogPosts, &blogPost)
	}

	if err := curr.Err(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(blogPosts)
}
