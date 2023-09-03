package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

/**
 * Lock pessimista - faz o lock da tabela, na linha, nesse momento ninguém pode alterar a linha
 * Lock otimista - faz o lock por versão, se alguém alterar a linha, a versão muda e o lock é feito na nova versão
 * Isso serve para problemas com concorrência
 * Aqui teremos que ter dado drop na tabela antes de executar o código
 */
 func main() {
	dsn := "root:root@tcp(192.168.157.51:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// O Gorm por padrão sempre usa transação, neste exemplo vamos criar transação manualmente, nesse exemplo vamos usar o lock pessimista
	tx := db.Begin()
	var category Category
	// No momento que fazemos o lock, ninguém pode alterar a linha, o debug é para mostrar o SQL que está sendo executado
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 1).Error
	if err != nil {
		panic(err)
	}

	category.Name = "Categoria 1 - Alterada"
	tx.Save(&category)
	// Aqui vamos fazer o commit da transação, após o commit o lock é liberado
	tx.Commit()


}
