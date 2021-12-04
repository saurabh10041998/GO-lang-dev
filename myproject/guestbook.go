package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"bufio"
	"fmt"
)


type Guestbook struct {
	SignatureCount int
	Signatures []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(filename string) [] string {
	var lines []string
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		line := Scanner.Text()
		lines = append(lines, line)
	}
	check(Scanner.Err())
	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	lines := getStrings("signature.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	data := Guestbook { SignatureCount: len(lines), Signatures: lines }
	err =  html.Execute(writer, data)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signature.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:9001", nil)
	log.Fatal(err)
}