package main

import (
	"os"
	"log"
	"text/template"
)


type Part struct { 
	Name string
	Count int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
} 

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	executeTemplate("dot is {{ . }}\n", "ABC")
	executeTemplate("dot is {{ . }}\n", 123.45)
	// ****** IF templates ********
	executeTemplate("start {{ if .}} dot is true!! {{ end }} finish\n", true)
	executeTemplate("start {{ if . }} dot is true!! {{ end }} finish\n", false)

	// ******* Range Template *********
	templateText :=  "[*] Before loop: {{ . }}\n {{ range .}}[+] In Loop {{ . }} \n {{ end }}[*]After loop: {{ . }}\n"
	executeTemplate(templateText, []string {"Sa", "Re", "Ga"})
	templateText = "[+] Prices: \n {{ range . }} $ {{ . }} \n {{ end }}"
	executeTemplate(templateText, []float64{23.45, 56.66, 78.98})
	executeTemplate(templateText, []float64{})
	executeTemplate(templateText, nil)

	// ******** Struct field **********
	templateText = "[*] Name: {{ .Name }}\n[*]Count: {{ .Count}}\n"
	executeTemplate(templateText, Part{ Name: "Saurabh", Count: 234 })
	executeTemplate(templateText, Part{ Name: "Abhishek", Count: 2344})
}