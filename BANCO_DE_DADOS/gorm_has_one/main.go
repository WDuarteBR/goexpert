package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Id   int `gorm:"primarykey"`
	Name string
}

type Product struct {
	Id           int `gorm:"primarykey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber //descreve uma relação "tem um"
	gorm.Model
}

type SerialNumber struct {
	Id        int `gorm:"primarykey"`
	Number    string
	ProductID int // também denota um relacionamento do tipo one to one
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})

	db.Create(&Category{
		Name: "Info",
	})

	db.Create(&Product{
		Name:       "Teclado Mecânico",
		Price:      139.98,
		CategoryID: 1,
	})

	db.Create(&SerialNumber{
		Number:    "xpto1234",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}

}
