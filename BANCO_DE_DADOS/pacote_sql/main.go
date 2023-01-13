package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}

}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("iPhone 13 pro", 7299.90)

	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Print("Cadastrado novo produto\n")

	product.Price = 8348.35
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := findProduct(db, product.Id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Atualizado : %s - no varlor de: %.2f\n", p.Name, p.Price)

	products, err := allProducts(db)
	if err != err {
		panic(err)

	}

	fmt.Println("##########TABELA DE PREÇOS##########")
	for _, p := range products {
		fmt.Printf("%s - %s -----------------%.2f\n", p.Id, p.Name, p.Price)
	}

	// id := "1624a512-57a4-47aa-a53b-bd21ed9efe10"

	err = deleteProduct(db, product.Id)
	if err != nil {
		panic(err)
	}

}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func allProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)

	}

	return products, nil
}

func findProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Id)
	if err != nil {
		return err
	}

	return nil

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?);")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Id, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}
