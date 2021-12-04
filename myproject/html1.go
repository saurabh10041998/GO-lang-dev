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
	text := "This is text to print\n"
	tmpl, err := template.New("Test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}