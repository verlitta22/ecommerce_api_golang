package models

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type Order struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `json:"user_id"`
	TotalPrice  float64     `json:"total_price"`
	CreatedAt   time.Time   `json:"created_at"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}