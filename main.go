package main

import (
	"ecommerce_api/config"
	"ecommerce_api/controllers"
	"ecommerce_api/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi Koneksi Database
	config.ConnectDB()

	// 1. Public Routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/products", controllers.GetProducts)

	// 2. Protected Routes (Membutuhkan Header "Authorization: Bearer <token>")
	api := app.Group("/api", middleware.Protected())
	api.Post("/products", controllers.CreateProduct)
	api.Post("/checkout", controllers.Checkout)

	app.Listen(":8080")
}