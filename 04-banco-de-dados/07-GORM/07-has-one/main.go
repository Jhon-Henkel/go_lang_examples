package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	Number    string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

// aqui teremos que ter dado drop na tabela antes de executar o código
func main() {
	dsn := "root:root@tcp(192.168.157.51:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Categoria 1"}
	db.Create(&category)

	product := Product{
		Name:       "Produto 1",
		Price:      10.0,
		CategoryID: category.ID,
	}
	db.Create(&product)

	serialNumber := SerialNumber{
		ProductID: product.ID,
		Number:    "123456789",
	}
	db.Create(&serialNumber)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
