package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Hello World From Hello Handler"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
