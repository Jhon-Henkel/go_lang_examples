package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blogHandler{title: "Blog"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World, mux2!"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello World, mux1!"))
}

type blogHandler struct {
	title string
}

func (blog blogHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(blog.title))
}
