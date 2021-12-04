package main

import (
	"log"
	"os"
	"text/template"
)

type Subscriber struct {
	Name string
	Active bool
	Rate float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout,data)
	check(err)
}


func main() {
	templateText := " [*] Name: {{ .Name }}\n{{ if .Active }} [*] Rate: {{ .Rate }} {{ end }}\n"
	executeTemplate(templateText, Subscriber{ Name: "John", Rate: 4.55, Active: true})
	executeTemplate(templateText, Subscriber{ Name: "Dave", Rate: 5.55, Active: false})
}