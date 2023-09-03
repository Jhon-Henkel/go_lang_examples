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

	// find all wirh limit
	var products []Product
	db.Limit(2).Find(&products)
	fmt.Println(products)

	// find all wirh limit and offset (pagination)
	var products1 []Product
	db.Limit(2).Offset(2).Find(&products1)
	fmt.Println(products1)

	// find with where
	var products2 []Product
	db.Where("price > ?", 10).Find(&products2)
	fmt.Println(products2)

	// find with like
	var products3 []Product
	db.Where("name LIKE ?", "%Produto%").Find(&products3)
	fmt.Println(products3)
}
