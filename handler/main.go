package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis/model"
	"go-redis/repository"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type Main struct {
	Cache      *redis.Client
	Repository repository.Main
}

func New(c *redis.Client, m repository.Main) Main {
	return Main{Cache: c, Repository: m}
}

func (handler Main) IndexProducts(ctx echo.Context) error {
	//cek dulu ke redis nya ada apa tidak?
	resultStr, err := handler.Cache.Get(context.Background(), "ftgo:demo_redis:products").Result()

	//kalo ada yg dibalikin dari redis
	if err == nil {
		result := []model.Product{}
		json.Unmarshal([]byte(resultStr), &result)
		return ctx.JSON(http.StatusOK, result)
	}

	//kalo gaada, baru ambil dari repo/service grpc/service http
	products := handler.Repository.FindAllProducts()

	//setelah diambil, save datanya ke redis
	productByte, err := json.Marshal(products)
	if err != nil {
		fmt.Println("failed to save products to redis", err.Error())
	} else {
		err = handler.Cache.Set(
			context.Background(), 
			"ftgo:demo_redis:products", 
			string(productByte), 
			10*time.Second,
		).Err()
		
		if err != nil {
			fmt.Println("failed to save products to redis", err.Error())
		}
	}

	//kasih response ke client
	return ctx.JSON(http.StatusOK, products)
}