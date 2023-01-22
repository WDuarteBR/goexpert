package main

import (
	"net/http"

	"github.com/go-chi/jwtauth"

	_ "github.com/wduartebr/goexpert/apis/docs"
	"github.com/wduartebr/goexpert/apis/internal/infra/database"
	"github.com/wduartebr/goexpert/apis/internal/infra/webserver/handlers"

	"github.com/wduartebr/goexpert/apis/internal/entity"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/wduartebr/goexpert/apis/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Expert API Exercise
// @version 1.0
// @description Product API with authentication
// @termOfService http://swagger.io/terms/

// @contact.name W. Duarte
// @contact.url http://aindafarei.com.br
// @contact.mail wduarte.br@gmail.com

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("experesin", config.JWTExperesIn))

	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB)

	r.Route("/product", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/all", productHandler.AllProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)

	})

	r.Post("/user", userHandler.CreateUser)
	r.Post("/user/gen_token", userHandler.GetJWT)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)

}
