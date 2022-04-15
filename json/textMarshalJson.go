package main;

import (
    "fmt"
    "log"
    "strings"
    "encoding/json"
)

type Size int

const (
    Unrecognized Size = iota
    Large
    Small
)

func (s *Size) UnmarshalText(text []byte) error {
    switch strings.ToLower(string(text)) {
        default:
           *s = Unrecognized
        case "small": 
           *s = Small
        case "large":
           *s = Large
    }
    return nil    
}

func (s Size) MarshalText()([]byte, error) {
    var text string
    switch s {
        default:
           text = "Unrecognized"
        case Small:
            text = "Small"
        case Large:
            text = "Large"

    }
    return []byte(text),nil
}


func main() {
    blob := `["small","regular","large","unrecognized","small","normal","small","large"]`

    var inventory []Size

    if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
            log.Fatal(err)
    }

    counts := make(map[Size]int)
    for _, token := range inventory {
        counts[token] += 1
    }

    fmt.Printf("[*] Inventory:\n \x1B[32m[*]Small -> %d\n \x1B[33m[*]Large -> %d\n \x1B[0m [*] Unrecognized -> %d\n", counts[Small], counts[Large], counts[Unrecognized]);

}
