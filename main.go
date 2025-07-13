package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"myProject/config"
	"myProject/handler"
	"myProject/router"
	"myProject/store"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment variables")
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

	storeSql, err := store.NewDBTable(dbPool)
	if err != nil {
		log.Fatalf("cant create tables: %v\n", err)
	}

	db := &handler.StoreStruct{DB: storeSql.DbPool}

	r := gin.Default()

	router.Router(r, db)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
