package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Itens struct {
	Id         int `gorm:"primarykey"`
	Name       string
	Price      float64
	gorm.Model // atribuirá colunas de controle com created_at
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Itens{})

	// adicionando um reqistro
	/*
		db.Create(&Itens{
			Name:  "TV Sansug 65\"",
			Price: 4658.90,
		})
	*/
	// criando em batch
	/*
		itens := []Itens{
			{Name: "Air Frier", Price: 657.99},
			{Name: "Xbox series X", Price: 4585.79},
			{Name: "Monitor wide screem LG", Price: 1138.99},
		}
		db.Create(itens)
	*/

	// var it Itens
	// selecionando um registro, pelo id
	//db.First(&it, 3)

	// selecionando um registro pelo name
	//db.First(&it, "name = ?", "Xbox series X")

	//selecionando todos os registros
	//var itens []Itens

	// limite e paginação
	//db.Limit(2).Offset(2).Find(&itens)

	//usando where
	//db.Where("price < ?", 2000).Find(&itens)
	// db.Where("name Like ?", "%box%").Find(&itens)

	// for _, item := range itens {

	// 	fmt.Println(item)
	// }

	//Alterando registro
	var it Itens
	// db.First(&it, 1)
	// it.Name = "Monitor WideScreem"
	// db.Save(&it)
	// fmt.Println("Registro alterado com sucesso")

	//Deletando registro

	db.First(&it, 2)
	db.Delete(&it)
	fmt.Println("Registro deletado")

	// fmt.Println(it)

}
