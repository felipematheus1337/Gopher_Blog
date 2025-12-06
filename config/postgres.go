package config

import (
	"fmt"
	"os"

	"github.com/felipematheus1337/GoPHER_Blog/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(postgres.Open(dbPath), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error initializing database: %v", err)
	}

	err = db.AutoMigrate(&schemas.Post{})

	if err != nil {
		return nil, fmt.Errorf("error creating the schema of the  database: %v", err)
	}

	return db, nil
}
