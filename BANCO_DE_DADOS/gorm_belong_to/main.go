package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categories struct {
	Id   int `gorm:"primarykey"`
	Name string
}

type Itens struct {
	Id         int `gorm:"primarykey"`
	Name       string
	Price      float64
	CategoryId int
	Category   Categories //descreve a relação "é um"
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Categories{}, &Itens{})

	/*
		category := Categories{
			Name: "Info",
		}
		db.Create(&category)

		item := Itens{
			Name:       "Teclado mecânico",
			Price:      158.98,
			CategoryId: category.Id,
		}
		db.Create(&item)
	*/

	var itens []Itens

	//Para retornar os relacionamentos faz-se da forma abaixo
	db.Preload("Category").Find(&itens)

	for _, item := range itens {
		fmt.Println(item)
	}

}
