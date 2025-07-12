package config

import (
	"log"
	"os"
	"strconv"
)

func GetAllFromEnv() (dbName, dbPassword, dbUser, dbHost string, dbPort uint16, err error) {
	dbName = os.Getenv("POSTGRES_DB")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbUser = os.Getenv("POSTGRES_USER")
	dbHost = os.Getenv("POSTGRES_HOST")
	dbPorStr := os.Getenv("POSTGRES_PORT")
	dbPortUint64, err := strconv.ParseUint(dbPorStr, 10, 16)

	if err != nil {
		log.Fatalf("Unable to parse database port or nan: %v\n", err)
		return
	}
	dbPort = uint16(dbPortUint64)
	return
}
