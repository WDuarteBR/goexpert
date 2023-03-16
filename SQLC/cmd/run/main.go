package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wduartebr/goexpert/sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbcnn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbcnn.Close()

	queries := db.New(dbcnn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Desc Backend", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "1f795b7d-3e52-4e1b-83dd-1284a27bece5",
		Description: sql.NullString{String: "Backend Desc Updated", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("*******LISTA DE CATEGORIAS*******")
	for _, category := range categories {
		fmt.Printf("Id: %s \nName: %s \nDescription: %s \n", category.ID, category.Name, category.Description.String)
		fmt.Println("------------------------------------")
	}
	fmt.Print("***********************************")

	err = queries.DeleteCategory(ctx, "1f795b7d-3e52-4e1b-83dd-1284a27bece5")

	if err != nil {
		panic(err)
	}

}
