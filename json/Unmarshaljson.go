package main

import  (
    "encoding/json"
    "fmt"
)

func main() {
    var jsonBlobs = []byte(`[
        {"name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll", "Order": "Dasyuromorphia"}
    ]`)

    type Animal struct {
        Name    string
        Order   string
    }

    var animals []Animal
    if err := json.Unmarshal(jsonBlobs, &animals); err != nil {
        fmt.Println("Error: ", err)
    }
    fmt.Printf("%+v", animals)
}
