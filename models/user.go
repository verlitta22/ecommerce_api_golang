package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `gorm:"default:'customer'" json:"role"` // 'admin' atau 'customer'
}