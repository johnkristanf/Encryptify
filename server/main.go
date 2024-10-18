package main

import (
	"net/http"

	"encryptify/database"
	"encryptify/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main(){
	e := echo.New()

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