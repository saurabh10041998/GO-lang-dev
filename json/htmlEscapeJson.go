package main;

import (
    "bytes"
    "os"
    "encoding/json"
)

func main() {
    var out bytes.Buffer
    blob := `{"name": "<b>HTML content</b>"}`
    json.HTMLEscape(&out, []byte(blob))
    out.WriteTo(os.Stdout)
}
