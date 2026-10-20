package controllers

import (
	"ecommerce_api/config"
	"ecommerce_api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// === PRODUCT CONTROLLER ===

// Get All Products (Public)
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	config.DB.Find(&products)
	return c.JSON(fiber.Map{"data": products})
}

// Create Product (Protected)
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Input tidak valid"})
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal menambah produk"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Produk berhasil ditambahkan", "data": product})
}

// === TRANSACTION / CHECKOUT CONTROLLER ===

type CheckoutItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CheckoutInput struct {
	Items []CheckoutItem `json:"items"`
}

func Checkout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input CheckoutInput
	if err := c.BodyParser(&input); err != nil || len(input.Items) == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Item pesanan tidak boleh kosong"})
	}

	// Gunakan DB Transaction agar konsisten
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var totalPrice float64 = 0
		var orderItems []models.OrderItem

		for _, item := range input.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				return fiber.NewError(404, "Produk tidak ditemukan")
			}

			if product.Stock < item.Quantity {
				return fiber.NewError(400, "Stok produk "+product.Name+" tidak mencukupi")
			}

			// Kurangi stok produk
			product.Stock -= item.Quantity
			if err := tx.Save(&product).Error; err != nil {
				return err
			}

			itemPrice := product.Price * float64(item.Quantity)
			totalPrice += itemPrice

			orderItems = append(orderItems, models.OrderItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     itemPrice,
			})
		}

		order := models.Order{
			UserID:     userID,
			TotalPrice: totalPrice,
			OrderItems: orderItems,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Checkout berhasil! Stok produk telah diperbarui."})
}

// === GET ORDER HISTORY CONTROLLER (BARU) ===
func GetOrderHistory(c *fiber.Ctx) error {
	// Ambil user_id dari JWT claims yang disimpan oleh middleware
	userID := c.Locals("user_id").(uint)

	var orders []models.Order
	// Ambil data order berdasarkan user_id, sekaligus preload (ambil data) item produk di dalamnya
	if err := config.DB.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil riwayat pesanan"})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil riwayat pesanan",
		"data":    orders,
	})
}