package routes

import (
	"log"

	"app/controllers"
	"app/controllers/blog"
	"app/controllers/products"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	router.Get("/:id", controllers.GetUser)
	router.Post("/sign-up", controllers.SignUp)
	router.Post("/login", controllers.LoginUser)
	router.Delete("/delete/:id", controllers.DeleteUser)

	log.Fatal("Error has occured")
}

func ReviewRoutes(router fiber.Router) {

}

func ProductRoutes(router fiber.Router) {
	router.Get("/:category", products.GetProductByCategory)

}

func BlogRoutes(router fiber.Router) {
	router.Get("/blog", blog.GetAllBlogPosts)
	router.Get("/blog/:id", blog.GetBlogPost)
	router.Post("/create-blog-post", blog.CreateBlogPost)
	router.Delete("/delete-blog/:id", blog.DeleteBlogPost)

	log.Fatal("error in blog routes")
}
