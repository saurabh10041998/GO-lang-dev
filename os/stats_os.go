package main;

import (
    "os"
    "log"
    "fmt"
)


func main() {
    /* get environment variable array */
    env_array := os.Environ()
    fmt.Println(env_array)

    /*get executable path*/
    exe_path,err := os.Executable()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\x1B[0m[*] Exe path: -> %s\n", exe_path)    
    /* set and get env vars */

    os.Setenv("NAME", "Asta")
    os.Setenv("SQUAD", "Black bulls")

    fmt.Printf("I am %s and I am from %s", os.Getenv("NAME"), os.Getenv("SQUAD"))

}
