package main;


import (
    "fmt"
    "encoding/json"
    "log"
    "strings"
)

type Animal int

const (
    Unknown Animal = iota           // what is this crap
    Gopher
    Zebra
)

func (a *Animal) UnmarshalJSON (b []byte) error {
     var s string
     if err := json.Unmarshal(b, &s); err != nil {
        return err
     }
     switch strings.ToLower(s) {
        default:
            *a =  Unknown
        case "gopher":
            *a = Gopher
         case "zebra":
            *a = Zebra    
     }
     return nil
}

func (a Animal) MarshalJSON()([]byte, error) {
    var s string
    switch a {
        case Unknown:
            s = "unknown"
        case Gopher:
            s = "Gopher"
        case Zebra:
            s = "Zebra"
    }
    return json.Marshal(s)
}

func main() {
    blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
    
    var zoo []Animal

    if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
        log.Fatal(err)
    }

    census := make(map[Animal]int)

    for _,animal := range zoo {
        census[animal] += 1
    }

    fmt.Printf("Zoo Census:\n\x1B[32m* Gophers: %d\n\x1B[33m* Zebras: %d\n\x1B[0m* Uknown: %d\n", census[Gopher], census[Zebra], census[Unknown])

    //Playing with the marshal 
    b, err := json.Marshal(&zoo);
    fmt.Println(string(b));
}


