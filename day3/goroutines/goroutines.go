package main

import (
    "fmt"
    "time"
)
//A goroutine is a lightweight thread of execution.

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")//Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.

    go f("goroutine")//To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.

    go func(msg string) {//You can also start a goroutine for an anonymous function call.
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)//Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
    fmt.Println("done")
}