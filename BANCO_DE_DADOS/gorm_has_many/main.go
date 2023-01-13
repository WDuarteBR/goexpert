package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Id       int `gorm:"primarykey"`
	Name     string
	Products []Product
}

type Product struct {
	Id           int `gorm:"primarykey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	Id        int `gorm:"primarykey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})
	/*
		db.Create(&Category{
			Name: "Escritorio",
		})

		db.Create(&Product{
			Name:       "Cadeira Presidente",
			Price:      879.98,
			CategoryID: 1,
		})

		db.Create(&SerialNumber{
			Number:    "789xyz",
			ProductID: 1,
		})
	*/
	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println("- ", p.Name, "- SN:", p.SerialNumber.Number)
		}
	}

}
