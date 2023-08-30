package main

import (
	"context"
	"database/sql"
	"fmt"

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
	procductSearch, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(procductSearch)
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

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectProductWithContext(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var product Product
	err = stmt.QueryRowContext(ctx, id).Scan(&product.ID, &product.Name, &product.price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
