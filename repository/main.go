package repository

import (
	"fmt"
	"go-redis/model"
	"time"
)

type Main struct{}

func New() Main {
	return Main{}
}

func (repository Main) FindAllProducts() []model.Product {
	//logic ke db
	time.Sleep(2 * time.Second)
	fmt.Println("querying to database...")
	return []model.Product{
		{Name: "Iphone 11", Price: 120},
		{Name: "Iphone 12", Price: 140},
		{Name: "Iphone 13", Price: 170},
		{Name: "Iphone 14", Price: 180},
		{Name: "Iphone 15", Price: 200},
	}
}
