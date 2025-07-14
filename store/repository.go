package store

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"myProject/models"
)

func CheckUserExist(username string, email string, db *pgxpool.Pool) (bool, error) {
	q := `SELECT * FROM users WHERE email=$1 or username=$2`
	find := &models.User{}
	err := db.QueryRow(context.Background(), q, username, email).Scan(find)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateUser(user *models.User, db *pgxpool.Pool) error {
	q := `INSERT INTO users (username, password_hashed, email, role) VALUES ($1, $2, $3, $4) RETURNING user_id, created_at`

	err := db.QueryRow(context.Background(), q, user.Username, user.Password, user.Email, user.Role).Scan(&user.UserID, &user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
