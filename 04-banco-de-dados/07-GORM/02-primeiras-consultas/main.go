package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(192.168.157.51:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
	db.Create(&Product{Name: "Produto 1", Price: 10.0})
	products := []Product{
		{Name: "Produto 2", Price: 20.0},
		{Name: "Produto 3", Price: 30.0},
	}
	db.Create(&products)

	// find by id
	var product Product
	db.First(&product, 1)
	fmt.Println(product)

	// find by name
	var product2 Product
	db.Find(&product2, "name = ?", "Produto 2")
	fmt.Println(product2)

	// find all
	var products2 []Product
	db.Find(&products2)
	fmt.Println(products2)
}
