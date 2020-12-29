package main
//We can use channels to synchronize execution across goroutines. 
//Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
import (
    "fmt"
    "time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	//This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true //Send a value to notify that we’re done.
}

func main() {

    done := make(chan bool, 1)//Start a worker goroutine, giving it the channel to notify on.


    go worker(done) 

    <-done // Block until we receive a notification from the worker on the channel.

}

//If you removed the <- done line from this program, the program would exit before the worker even started.

