package main
//We often want to execute Go code at some point in the future, or repeatedly at some interval.
// Go’s built-in timer and ticker features make both of these tasks easy. We’ll look first at timers and then at tickers.
import (
    "fmt"
    "time"
)

func main() {

    timer1 := time.NewTimer(2 * time.Second)
	//Timers represent a single event in the future. 
	//You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
    <-timer1.C
	//The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
    fmt.Println("Timer 1 fired")
	timer2 := time.NewTimer(time.Second)
	//If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

	time.Sleep(2 * time.Second)
	//Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
}