package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_"github.com/lib/pq"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	// log.Println("DB_PORT:", os.Getenv("DB_PORT"))
	// log.Println("DB_USER:", os.Getenv("DB_USER"))
	// log.Println("DB_NAME:", os.Getenv("DB_NAME"))
	// log.Println("DB_SSL_MODE:", os.Getenv("DB_SSL_MODE"))
	// log.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL_MODE"),
		)
	}

	log.Println("Connecting with DSN:", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to the PostgreSQL database successfully")
	return db, nil
}
