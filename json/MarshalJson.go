package main

import (
    "encoding/json"
    "os"
    "fmt"
)

func main() {
    type ColorGroup struct {
        ID      int     `json:"customId"`
        Name    string  `json:"grpName,omitempty"`
        Colors  []string 
        Ignore  bool    `json:"-"`
    }

    group := ColorGroup { 
        ID:     1,
        Name:   "Reds",
        Colors: []string{"Crimson", "Black", "Red", "Ruby", "Maroon"},
        Ignore: false,
    }
    b,err := json.Marshal(group)
    if err != nil { 
        fmt.Println("Error :", err)
    }
    os.Stdout.Write(b)   
    
    group2 := ColorGroup {
        ID:     2,
        Name: "",
        Colors: []string{"Black", "Gold", "Silver"},
        Ignore: true,
    }
    b2, err := json.Marshal(group2)
    if err != nil {
        fmt.Println("Error 2", err)
    }
    fmt.Println()           //  insert the enter :) 
    os.Stdout.Write(b2)
} 
