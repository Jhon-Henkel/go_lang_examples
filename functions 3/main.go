package main

func main() {
	sum := func() int {
		return sumNumbers(5, 6, 7, 4) * 2
	}()
	printSum(sum)
}

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func printSum(sum int) {
	println(sum)
}
