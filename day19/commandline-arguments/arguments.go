package main

import (
    "fmt"
    "os"
)

func main() {

    argsWithProg := os.Args //os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3] //You can get individual args with normal indexing.

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}