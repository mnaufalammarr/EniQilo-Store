package db

import (
	"EniQilo/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
	"os"
)

func InitDB() *pgxpool.Pool {

	utils.LoadEnvVariables()
	// urlDb := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	// dbParams := os.Getenv("DB_PARAMS")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	conn, err := ConnectToDatabase(dbURL)

	if err != nil {
		log.Fatal("db connection failed")
	}

	return conn
}

func ConnectToDatabase(urlDb string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), urlDb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return dbpool, err
}

func ClosePool(dbPool *pgxpool.Pool) {
	defer dbPool.Close()
}
