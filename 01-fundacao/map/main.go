package main

import "fmt"

func main() {
	salarios := map[string]int{"mínimo": 1320, "recebo": 3500, "diferença": 2180}

	fmt.Println("salário Mínimo:", salarios["mínimo"])

	delete(salarios, "mínimo")
	salarios["teste"] = 1000

	fmt.Println("salário Mínimo:", salarios["mínimo"])
	fmt.Println("salário Mínimo:", salarios["teste"])

	sal := make(map[string]int)

	fmt.Println(sal)

	for typeName, value := range salarios {
		fmt.Println(typeName, value)
	}

	// o underline é um blank identifier, ou seja, ele ignora o primeiro valor
	for _, value := range salarios {
		fmt.Println(value)
	}
}
