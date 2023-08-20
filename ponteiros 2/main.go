package main

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	numberOne := 10
	numberTwo := 20
	total := sum(&numberOne, &numberTwo)
	println(total, numberOne, numberTwo)
}
