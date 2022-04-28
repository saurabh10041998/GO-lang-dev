package main

import (
	"log"
	"os"
)
func main() {
	err := os.Mkdir("testdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("testdir/testfile.txt", []byte("hello Gophers!"),0660)
	if err != nil {
		log.Fatal(err)		
	}
	log.Println("[+] How does it look ??")
}
