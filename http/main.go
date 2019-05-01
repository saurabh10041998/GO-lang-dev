package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// io.Copy(os.Stdout, resp.Body)

	//Writing custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("Just wrote these many bytes:", len(p))
	return len(p), nil
}
