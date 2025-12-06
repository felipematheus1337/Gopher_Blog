package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felipematheus1337/GoPHER_Blog/config"
	handler2 "github.com/felipematheus1337/GoPHER_Blog/handler"
	"github.com/felipematheus1337/GoPHER_Blog/router"
	"github.com/felipematheus1337/GoPHER_Blog/service"
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

	db, err := config.InitializePostgres()

	if err != nil {
		log.Fatalf("Error initializing postgres : %v", err)
	}

	services := service.NewPostService(db)

	handler := handler2.NewPostHandler(services)

	router.InitializeRoutes(r, handler)

	r.Run(port)

}
