package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	json := bytes.NewBuffer([]byte(`{"nome":"joao"}`))
	res, err := client.Post("https://google.com", "application/json", json)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}
