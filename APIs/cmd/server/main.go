package main

import (
	"net/http"

	"github.com/wduartebr/goexpert/apis/internal/infra/database"
	"github.com/wduartebr/goexpert/apis/internal/infra/webserver/handlers"

	"github.com/wduartebr/goexpert/apis/internal/entity"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/wduartebr/goexpert/apis/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	productHandler := handlers.NewProductHandler(productDB)

	r.Post("/product", productHandler.CreateProduct)
	r.Get("/product/{id}", productHandler.GetProduct)
	r.Put("/product/{id}", productHandler.UpdateProduct)
	http.ListenAndServe(":8000", r)

	// 5b5f540a-3b11-4f70-9ec2-b9579a1d61c2
}
