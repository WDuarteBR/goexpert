package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Id       int `gorm:"primarykey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}

type Product struct {
	Id         int `gorm:"primarykey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	/*
		db.AutoMigrate(&Category{}, &Product{})

		cat1 := Category{
			Name: "Escritorio",
		}
		db.Create(&cat1)

		cat2 := Category{
			Name: "Info",
		}
		db.Create(&cat2)

		db.Create(&Product{
			Name:       "Suporte para NoteBook",
			Price:      100.00,
			Categories: []Category{cat1, cat2},
		})
	*/

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println("- ", p.Name)
		}
	}

}
