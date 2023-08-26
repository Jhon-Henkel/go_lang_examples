package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria string
}

type Cursos []Curso

func main() {
	templateMust := template.Must(template.New("template.html").ParseFiles("template.html"))
	cursos := Cursos{
		Curso{"Curso de Go", "40h"},
		Curso{"Curso de Python", "30h"},
		Curso{"Curso de JavaScript", "20h"},
	}
	err := templateMust.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}
