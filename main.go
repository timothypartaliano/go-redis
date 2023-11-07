package main

import (
	"go-redis/config"
	"go-redis/handler"
	"go-redis/repository"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	repo := repository.New()
	cache := config.InitCache()

	mainHandler := handler.New(cache, repo)

	app.GET("/products", mainHandler.IndexProducts)

	log.Fatal(app.Start(":8080"))
}