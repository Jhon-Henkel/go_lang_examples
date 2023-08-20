package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// gravar
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	//tam, err := f.WriteString("Hello, World!")
	tam, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}
	println(tam, "bytes escrito")
	f.Close()

	// ler (esse vai carregar o arquivo inteiro na memória)
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	println(string(file))

	// para ler o arquivo sem carregar inteiro na memória de uma vez, fazemos a leitura de pouco em pouco abrindo o arquivo
	file2, arr := os.Open("test.txt")
	if arr != nil {
		panic(arr)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 5)
	for {
		tam, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:tam]))
	}
	file2.Close()

	// remover
	err2 := os.Remove("test.txt")
	if err2 != nil {
		panic(err2)
	}
}
