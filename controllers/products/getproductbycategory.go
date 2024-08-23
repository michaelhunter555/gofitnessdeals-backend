package products

import (
	"context"
	"net/http"

	amazon "app/amazonclient"

	"github.com/goark/pa-api/entity"
	"github.com/goark/pa-api/query"
	"github.com/gofiber/fiber/v2"
)

func GetProductByCategory(c *fiber.Ctx) error {
	category := c.Params("category")

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	client := amazon.Client()

	q := query.NewSearchItems(
		client.Marketplace(),
		client.PartnerTag(),
		client.PartnerType(),
	).RequestFilters(query.RequestMap{
		query.Keywords:  category,
		query.ItemPage:  page,
		query.ItemCount: limit,
	})

	body, err := client.RequestContext(context.Background(), q)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No products found for the query"})
	}
	res, err := entity.DecodeResponse(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error Decoding Response body"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"products": res})
}
