package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	numeros := []string{"um", "dois", "tres"}
	for key, value := range numeros {
		println(key, value)
	}
	i := 0
	for i < 5 {
		println(i)
		i++
	}
}
