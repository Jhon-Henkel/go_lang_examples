package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria string
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"template.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templateNew := template.New("template.html")
		templateNew.Funcs(template.FuncMap{"ToUpper": ToUpper})
		templateMust := template.Must(templateNew.ParseFiles(templates...))
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
