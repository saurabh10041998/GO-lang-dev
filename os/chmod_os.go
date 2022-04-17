package main

import (
    "os"
    "fmt"
)

func main() {
    if err := os.Chmod("test.txt", 0777); err != nil {
        fmt.Println("[!!] Error :", err)
    }
    fmt.Println("Is it succed ??")
}
