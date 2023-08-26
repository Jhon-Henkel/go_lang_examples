package main

type MyNumber int

type Number interface {
	~int | float64
}

func soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T Number](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	m2 := map[string]float64{
		"key1": 1.1,
		"key2": 2.2,
		"key3": 3.3,
	}

	m3 := map[string]MyNumber{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	println(soma(m))
	println(soma(m2))
	println(soma(m3))
	println(compara(10, 10.0))
	println(compara(11, 10.0))
}
