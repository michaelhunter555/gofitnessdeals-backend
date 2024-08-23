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

	app.Use("/user", routes.UserRoutes)
	app.Use("/blog", routes.BlogRoutes)
	app.Use("/products", routes.ProductRoutes)

	log.Fatal(app.Listen(":5000"))
}
