package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("http://google.com.br")
	if err != nil {
		panic(err)
	}
	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(response))
	request.Body.Close()
}
