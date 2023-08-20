package main

import "fmt"

func main() {
	slice := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("tamanho = %d ; capacidade = %d ; valor %v\n", len(slice), cap(slice), slice)

	// esses dois pontos no array faz o drop de tudo que tiver a direita pa posição, apresentando somente o que tiver a esquerda
	fmt.Printf("tamanho = %d ; capacidade = %d ; valor %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])
	fmt.Printf("tamanho = %d ; capacidade = %d ; valor %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	// nesse caso ele faz o drop de tudo que tiver a esquerda da posição, apresentando somente o que tiver a direita
	fmt.Printf("tamanho = %d ; capacidade = %d ; valor %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// adicionando mais uma posição no slice, a capacidade do slice dobra
	slice = append(slice, 9)
	fmt.Printf("tamanho = %d ; capacidade = %d ; valor %v\n", len(slice), cap(slice), slice)
}
