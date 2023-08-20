package main

func main() {
	a := 10
	// esse & é para mostrar o endereço de memória
	println(&a)

	var pointer *int = &a
	println(pointer)
	*pointer = 20
	println(a)

	b := &a
	println(*b)
}
