package models

import (
	"time"
)

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password_hashed"` // پسورد هرگز در JSON نمی‌آید
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Product struct {
	ProductId   int       `json:"product_id"`
	Quantity    int       `json:"quantity"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	ProductName string    `json:"product_name"`
	Rate        int       `json:"rate"`
	Colors      []string  `json:"colors"`
	Labels      [4]string `json:"labels"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}

type Comment struct {
	CommentId int       `json:"comment_id"`
	Content   string    `json:"content"`
	UserId    int       `json:"user_id"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}
