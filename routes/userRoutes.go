package routes

import (
	"os"

	"app/controllers"
	"app/controllers/blog"
	"app/controllers/products"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func UserRoutes(router fiber.Router) {
	router.Get("/:id", controllers.GetUser)
	router.Post("/sign-up", controllers.SignUp)
	router.Post("/login", controllers.LoginUser)
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
	}))
	router.Use(middleware.ExtendSession)
	router.Delete("/delete/:id", controllers.DeleteUser)
}

func ReviewRoutes(router fiber.Router) {

}

func ProductRoutes(router fiber.Router) {
	router.Get("/:category", products.GetProductByCategory)
	router.Use(middleware.ExtendSession)

}

func BlogRoutes(router fiber.Router) {
	router.Get("/blog", blog.GetAllBlogPosts)
	router.Get("/blog/:id", blog.GetBlogPost)
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
	}))
	router.Use(middleware.ExtendSession)
	router.Post("/create-blog-post", blog.CreateBlogPost)
	router.Delete("/delete-blog/:id", blog.DeleteBlogPost)

}
