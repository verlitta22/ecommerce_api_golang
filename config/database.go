package config

import (
	"fmt"

	"ecommerce_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Tambahkan parameter parseTime, loc, dan clear text/TLS setting jika diperlukan
	dsn := "root:@tcp(127.0.0.1:3306)/db_ecommerce_golang?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Gagal terkoneksi ke database Laragon! Error: %v", err))
	}

	// Auto Migrate Tabel (Akan otomatis membuat tabel users, products, orders, dll)
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{})

	DB = db
	fmt.Println("Database Laragon Terkoneksi & Migrasi Berhasil!")
}