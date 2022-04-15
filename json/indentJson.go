package main

import (
    "bytes"
    "encoding/json"
    "os"
    "log"
)

func main() {
    type Road struct {
        Name string
        Number int
    }

    roads := []Road {
        {
            "Diamond fork",
            29,
        },
        {
            "Sheep Creek",
            51,
        },
        {
            "Koliwada",
            34,
        },
    }
    b, err := json.Marshal(roads)
    if err != nil {
        log.Fatal(err)
    }
    var out bytes.Buffer
    json.Indent(&out, b, "=", "\t")
    out.WriteTo(os.Stdout)
}
