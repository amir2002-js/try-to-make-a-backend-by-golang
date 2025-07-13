package models

import (
	"time"
)

type User struct {
	Username       string    `json:"username"`
	PasswordHashed string    `json:"password_hashed"`
	Email          string    `json:"email"`
	UserId         int       `json:"user_id"`
	RoleUser       string    `json:"role"`
	CreateAt       time.Time `json:"create_at"`
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
