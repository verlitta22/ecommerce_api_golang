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

	// 1. Public Routes (Bisa diakses tanpa login)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/products", controllers.GetProducts)
	app.Get("/api/categories", controllers.GetCategories) // Rute baru untuk melihat kategori & produknya

	// 2. Protected Routes (Membutuhkan Header "Authorization: Bearer <token>")
	api := app.Group("/api", middleware.Protected())
	api.Post("/products", controllers.CreateProduct)
	api.Post("/categories", controllers.CreateCategory)   // Rute baru untuk menambah kategori
	api.Post("/checkout", controllers.Checkout)
	api.Get("/orders", controllers.GetOrderHistory)

	app.Listen(":8080")
}