package main

import (
	"net/http"
	"log"
	
	"encryptify/database"
	"encryptify/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
)


func main(){
	e := echo.New()

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, // Allowed HTTP methods
	}))

	e.POST("/api/register", handlers.RegisterHandler)
	e.POST("/api/login", handlers.LoginHandler)
	e.GET("/api/users", handlers.FetchUsersHandler)

	if err := database.MigrateUsersTable(); err != nil{
		e.Logger.Printf("Error in Migrating Users Table: %v", err)
	}

	e.Logger.Fatal(e.Start(":8080"))

}