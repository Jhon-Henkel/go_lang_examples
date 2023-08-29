package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(nil)
	}
	println(string(body))
}
