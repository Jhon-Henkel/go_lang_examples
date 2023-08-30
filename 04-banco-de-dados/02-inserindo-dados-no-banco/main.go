package main

import (
	"database/sql"
	// Esse _ é necessário para que o go não reclame que não estamos usando o pacote e não remova ele
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.157.51:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Produto 1", 10.0)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.price)
	if err != nil {
		return err
	}
	return nil
}
