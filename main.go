package main

import (
	"log"

	"app/controllers"
	"app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	client := ConnectDB()

	userCollection := GetCollection(client, "user")
	blogCollection := GetCollection(client, "blog")

	controllers.SetUserController(userCollection)
	controllers.SetUserController(blogCollection)

	userGroup := app.Group("/user")
	blogGroup := app.Group("/blog")
	productGroup := app.Group("/products")

	routes.UserRoutes(userGroup)
	routes.BlogRoutes(blogGroup)
	routes.ProductRoutes(productGroup)

	log.Fatal(app.Listen(":5000"))
}
