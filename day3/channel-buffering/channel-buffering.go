package main
//By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values without a corresponding receiver for those values.
import "fmt"

func main() {

    messages := make(chan string, 2)//Here we make a channel of strings buffering up to 2 values.


    messages <- "buffered"//Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
    messages <- "channel"
	// messages <- "channel2" when you have a 2 string buffer if you try to fire 3 channels it crashes the program with a fatal error
	//fatal error: all goroutines are asleep - deadlock!

	//goroutine 1 [chan send]:
	//main.main()
	//        /Users/eren/Documents/OtherProg/Go/GoIn30/day3/channel-buffering/channel-buffering.go:13 +0x8d
	//exit status 2
	//Erens-Air:day3 eren$ 

    fmt.Println(<-messages)//Later we can receive these two values as usual.
    fmt.Println(<-messages)
}