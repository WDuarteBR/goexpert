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
	http.ListenAndServe(":8000", r)

}
