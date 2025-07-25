package store

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	DbPool *pgxpool.Pool
}

func NewDBTable(dbPool *pgxpool.Pool) (*DB, error) {
	queries := []string{
		createUserTable(),
		createProductTable(),
		createCommentTable(),
	}

	for _, query := range queries {
		_, err := dbPool.Exec(context.Background(), query)
		if err != nil {
			return nil, err
		}
	}

	return &DB{
		DbPool: dbPool,
	}, nil
}

func createProductTable() string {
	return `    CREATE TABLE IF NOT EXISTS products (
        product_id      SERIAL PRIMARY KEY,
        product_name    VARCHAR(155) NOT NULL,
        description     TEXT NOT NULL,
        price           INTEGER NOT NULL,
        quantity        INTEGER DEFAULT 0,
        rating          INTEGER DEFAULT 0,
        image_url       TEXT NOT NULL,
        colors          TEXT[] NOT NULL,
        labels          TEXT[] NOT NULL,
        created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    )`
}

func createUserTable() string {
	return `
    CREATE TABLE IF NOT EXISTS users (
        user_id         SERIAL PRIMARY KEY,
        username        VARCHAR(50) UNIQUE NOT NULL,
        password_hashed TEXT NOT NULL,
        email           TEXT UNIQUE NOT NULL,
        role            TEXT NOT NULL,
        created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    )`
}

func createCommentTable() string {
	return ` CREATE TABLE IF NOT EXISTS comments (
        comment_id  SERIAL PRIMARY KEY,
        content     TEXT NOT NULL,
        user_id     INTEGER NOT NULL,
        product_id  INTEGER NOT NULL,
        created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

        CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id),
        CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(product_id)
    )`
}
