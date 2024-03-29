package main;

import (
    "os"
    "log"
    "fmt"
)


func main() {
    
    var GREEN string = "\x1B[32m"
    var RESET string = "\x1B[0m"
    var YELLOW string = "\x1B[33m"

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

    fmt.Printf("I am %s and I am from %s\n", os.Getenv("NAME"), os.Getenv("SQUAD"))

    /* Effective Group id */
    egid := os.Getegid()

    fmt.Printf("%s[*] Effective Group ID: %d\n%s", GREEN, egid, RESET)
    
    /* Effective user id */
    euid := os.Geteuid()

    fmt.Printf("%s[*] Effective user ID: %d%s",GREEN, euid, RESET)

    /* Group id */
    gid := os.Getgid()
    
    fmt.Printf("\n%s[*] Group id: %d%s", GREEN, gid, RESET)

    /* Groups to which user belngs*/
    fmt.Printf("\n%s[+] Groups : \n%s", YELLOW, RESET)
    grps, err := os.Getgroups()

    if err != nil {
        log.Fatal(err)
    }
    for _, x := range grps {
        fmt.Println(x)
    }

    /* Getting the page size */
    pgSz := os.Getpagesize()

    fmt.Printf("\n%s[*] Page size: %d%s", GREEN, pgSz, RESET)

    /* Proc id */
    pid := os.Getpid()
    fmt.Printf("\n%s[*] Process id: %d%s", GREEN, pid, RESET)

    /* Getting parent proc id */
    ppid := os.Getppid()

    fmt.Printf("\n%s[*] Parent process id %d%s", GREEN, ppid, RESET)
    
    /*  Current working directory */
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\n%s[*] Current working directory: %s%s", GREEN, cwd, RESET)
}
