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

	// edit register
	var product Product
	db.First(&product, 1)
	product.Name = "New Produto 1"
	db.Save(&product)

	// delete register
	var product2 Product
	db.Find(&product2, "name LIKE ?", "%New%")
	fmt.Println(product2)
	db.Delete(&product2)
}
