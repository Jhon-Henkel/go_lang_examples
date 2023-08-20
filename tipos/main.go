package main

import (
	"fmt"
)

const a = "Hello, World!"

type id int

var (
	b bool    = true
	c string  = "aaa"
	d float32 = 1.2
	f id      = 1
)

func main() {
	jhon := "foda"
	jhon += " pra crl"
	fmt.Printf("O tipo de jhon é %T", jhon)
	fmt.Printf("O valor de c é %v", c)
}
