package main
//To wait for multiple goroutines to finish, we can use a wait group.
import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
	//This is the function we’ll run in every goroutine. Note that a WaitGroup must be passed to functions by pointer.

	defer wg.Done()
	//On return, notify the WaitGroup that we’re done.

    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second) //Sleep to simulate an expensive task.

    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
}