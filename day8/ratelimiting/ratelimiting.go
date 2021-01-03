package main
//Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.
import (
    "fmt"
    "time"
)

func main() {

	requests := make(chan int, 5)
	//First we’ll look at basic rate limiting. Suppose we want to limit our handling of incoming requests. We’ll serve these requests off a channel of the same name.
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	//This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.

    for req := range requests {
		//By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
        <-limiter
        fmt.Println("request", req, time.Now())
    }

	burstyLimiter := make(chan time.Time, 3)
	//We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. We can accomplish this by buffering our limiter channel. This burstyLimiter channel will allow bursts of up to 3 events.

    for i := 0; i < 3; i++ {
		//Fill up the channel to represent allowed bursting.
        burstyLimiter <- time.Now()
    }

    go func() {
		//Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

	burstyRequests := make(chan int, 5)
	//Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstyLimiter.
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("bursty request", req, time.Now())
    }
}