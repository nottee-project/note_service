package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/nottee-project/task_service/internal/config"

	router "github.com/nottee-project/task_service/internal/delivery/rest"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}
	log.Printf("Current working directory: %s", dir)

	cfg := &config.Config{
		AuthServiceURL: os.Getenv("AUTH_SERVICE_URL"),
	}

	if cfg.AuthServiceURL == "" {
		log.Fatal("AUTH_SERVICE_URL is not set")
	}

	e := echo.New()

	if err := router.RegisterRoutes(e, cfg.AuthServiceURL); err != nil {
		e.Logger.Fatalf("Error registering routes: %v", err)
	}

	for _, route := range e.Routes() {
		fmt.Printf("Route: %s %s\n", route.Method, route.Path)
	}

	fmt.Println("Server started at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
