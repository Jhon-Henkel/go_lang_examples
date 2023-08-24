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
	templateMust := template.Must(template.New("cursoTemplate").Parse("Curso: {{.Nome}}\nCarga Horária: {{.CargaHoraria}}\n"))
	err := templateMust.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
