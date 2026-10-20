package controllers

import (
	"ecommerce_api/config"
	"ecommerce_api/models"

	"github.com/gofiber/fiber/v2"
)

// Tambah Kategori Baru (Admin)
func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Input tidak valid"})
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal membuat kategori"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Kategori berhasil dibuat",
		"data":    category,
	})
}

// Lihat Semua Kategori Beserta Produknya
func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	// Preload digunakan untuk mengambil data produk yang masuk ke dalam kategori tersebut
	if err := config.DB.Preload("Products").Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data kategori"})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data kategori",
		"data":    categories,
	})
}