package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    goodJSON := `{"example": 1}`
    badJSON := `{"example": 2 ]}}}`

    fmt.Println(json.Valid([]byte(goodJSON)))
    fmt.Println(json.Valid([]byte(badJSON)))
}
