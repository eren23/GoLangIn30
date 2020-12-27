package main

import "fmt"

func main() {

    if 7%2 == 0 {// basic if else
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 { //if w/o else
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 { // parentheses are not but the braces are required around conditions
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

//There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.