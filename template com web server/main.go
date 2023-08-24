package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria string
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templateMust := template.Must(template.New("template.html").ParseFiles("template.html"))
		cursos := Cursos{
			Curso{"Curso de Go", "40h"},
			Curso{"Curso de Python", "30h"},
			Curso{"Curso de JavaScript", "20h"},
		}
		err := templateMust.Execute(w, cursos)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
