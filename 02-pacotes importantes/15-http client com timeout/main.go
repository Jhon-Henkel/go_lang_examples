package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: 2 * time.Microsecond}
	res, err := client.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bod, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(bod))
}
