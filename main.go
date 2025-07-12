package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"myProject/config"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dbName, dbPassword, dbUser, dbHost, dbPort, err := config.GetAllFromEnv()
	if err != nil {
		log.Fatalf("Unable to parse database port or nan: %v\n", err)
		return
	}

	conf, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v\n", err)
		return
	}

	setter := config.ConfigSet{Port: dbPort, User: dbUser, Host: dbHost, Password: dbPassword, DBName: dbName}
	dbPool, err := setter.SetConf(conf)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return
	}
	defer dbPool.Close()

	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		_ = fmt.Sprint("Hello World")
		c.JSON(200, gin.H{"message": "Hello World"})
	})
	r.GET("/user", func(c *gin.Context) {
		_ = fmt.Sprint("Hello ali")
		c.JSON(200, gin.H{"message": "Hi ali"})
	})

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
