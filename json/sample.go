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
    

}


