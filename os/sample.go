package main

import (
    "os"
    "log"
    "fmt"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    buffer := make([]byte, 100)
    count, err := file.Read(buffer) 
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("[*] Read %d bytes: %s\n", count, buffer[:count])
}
