package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil { 
		log.Fatal(err)
	}
}

func main() {
	text := "Template start\nAction: {{ . }}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 42)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
}