package main

import "fmt"

func main() {

    messages := make(chan string)//Create a new channel with make(chan val-type). Channels are typed by the values they convey.

    go func() { messages <- "ping" }() // Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.

    msg := <-messages //The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
    fmt.Println(msg)
}

//By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.