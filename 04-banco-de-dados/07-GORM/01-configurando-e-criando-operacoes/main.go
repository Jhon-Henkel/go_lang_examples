package main

import (
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
}
