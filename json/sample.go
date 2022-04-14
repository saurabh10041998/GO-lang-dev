package main;


import (
    "fmt"
    "encoding/json"
    "logs"
    "strings"
)

Type Animal int

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
     return nil;
}




