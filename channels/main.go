package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://udemy.com",
	}
	//making channel of type string for communication
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// recieving msg for all go routines
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// this piece of code executing concurrently and comm via Channels
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
	}
	fmt.Println(link, "is up!")
	c <- link
}
