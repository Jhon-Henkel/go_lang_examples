package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

// aqui teremos que ter dado drop na tabela antes de executar o código
func main() {
	dsn := "root:root@tcp(192.168.157.51:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}

	db.Create(&Product{
		Name:  "Produto 1",
		Price: 10.0,
	})

	var product Product
	db.First(&product, 1)
	product.Name = "New Produto 1"
	db.Save(&product)

	db.Delete(&product)
}
