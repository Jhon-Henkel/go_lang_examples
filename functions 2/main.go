package main

func main() {
	sum := sumNumbers(1, 2)
	printSum(sum)
	sum2 := sumNumbers(5, 6, 7, 4)
	printSum(sum2)
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
