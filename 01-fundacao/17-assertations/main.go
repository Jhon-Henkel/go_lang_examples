package main

import "fmt"

func main() {
	var myVar interface{} = "Hello"
	fmt.Println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Println(res, ok)

}
