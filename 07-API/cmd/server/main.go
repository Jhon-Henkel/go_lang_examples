package main

import "github.com/Jhon-Henkel/go_lang_examples/tree/main/07-API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DbDriver)
}