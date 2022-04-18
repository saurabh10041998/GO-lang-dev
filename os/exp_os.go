package main

import (
   "os"
   "fmt"
)

func main() {
    /* os.Expand demo */
    mapper := func (placeholder string) string {
        switch placeholder {
            case "DAY_PART":
                return "morning"
            case "NAME":
                return "asta"
        }
        return ""
    }

    fmt.Println(os.Expand("Good ${DAY_PART}, $NAME !!", mapper))
}
