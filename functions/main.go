package main

import "errors"

func main() {
	sum, err := isSumMoreBiggerThanTen(1, 2)
	printSum(sum, err)
	sum2, err2 := isSumMoreBiggerThanTen(5, 6)
	printSum(sum2, err2)
}

func isSumMoreBiggerThanTen(value1, value2 int) (int, error) {
	sum := value1 + value2
	if sum >= 10 {
		return sum, nil
	}
	return sum, errors.New("sum is less than 10")
}

func printSum(sum int, err error) {
	if err != nil {
		println(err.Error())
		return
	}
	println(sum)
}
