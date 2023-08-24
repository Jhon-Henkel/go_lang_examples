package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria string
}

func main() {
	curso := Curso{"Go Expert", "40h"}
	tmp := template.New("cursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}}\nCarga Horária: {{.CargaHoraria}}\n")
	err := tmp.ExecuteTemplate(os.Stdout, "cursoTemplate", curso)
	if err != nil {
		panic(err)
	}
}
