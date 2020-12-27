package main

import "fmt"

//for is Go’s only looping construct. Here are some basic types of for loops.

func main() {

    i := 1
    for i <= 3 { // Basic type one condition
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ { //A classic initial/condition/after for loop.
        fmt.Println(j)
    }

    for { //for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ { //You can also continue to the next iteration of the loop.
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}