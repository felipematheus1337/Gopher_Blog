package config

import (
	"fmt"
	"os"

	"github.com/felipematheus1337/GoPHER_Blog/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	err = db.AutoMigrate(&schemas.Post{})

	if err != nil {
		return nil, fmt.Errorf("Error auto-migrating database: %v", err)
	}

	return db, nil

}
