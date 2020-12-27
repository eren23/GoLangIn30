package main

import "fmt"

func main() {

    var a = "initial" // to declare one variable
    fmt.Println(a)

    var b, c int = 1, 2 //  multiple variable declaretion 
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int // redeclaring the same variable is not possible
    fmt.Println(e)

    var g string //while e above prints 0 unassigned string declaretion prints an empty string
    fmt.Println(g)

    f := "apple" // := for declaring and initilazing a variable, this line is the same as var f string = "apple"
    fmt.Println(f)
}