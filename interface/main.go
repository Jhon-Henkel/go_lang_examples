package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
}

type Person interface {
	disable()
}

type Cliente struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (e Enteprise) disable() {
}

func Disabled(person Person) {
	person.disable()
}

type Enteprise struct {
	Name string
}

func (c Cliente) disable() {
	c.Active = false
	fmt.Printf("O cliente %s agora está desativado", c.Name)
}

func main() {
	jhon := Cliente{
		Name:   "Jhon",
		Age:    20,
		Active: true,
	}
	jhon.Street = "Rua dos bobos"
	jhon.Number = 10
	jhon.City = "São Paulo"

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", jhon.Name, jhon.Age, jhon.Active)
	fmt.Printf("Rua: %s\nNúmero: %d\nCidade: %s\n", jhon.Address.Street, jhon.Address.Number, jhon.Address.City)

	jhon.disable()
	Disabled(jhon)
}
