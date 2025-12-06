package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felipematheus1337/GoPHER_Blog/config"
	"github.com/gin-gonic/gin"
)

func main() {

	err := config.Init()

	if err != nil {
		fmt.Println("Error loading initializing...")
	}

	r := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	r.Run(":" + port)

	db, err := config.InitializePostgres()

	if err != nil {
		log.Fatalf("Error initializing postgres : %v", err)
	}

	_ = db

}
