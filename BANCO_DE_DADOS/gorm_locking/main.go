package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	db.AutoMigrate(&Category{}, &Product{})

	/*
		Lock otimista:
		Quando à um versionamento do registro

		Lock persimista:
		Quando o registro fica travado até que as alterações sejam comitadas
	*/

	tx := db.Begin() // transação criada
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}

	c.Name = "Office"
	db.Debug().Save(&c)
	tx.Commit()
}
