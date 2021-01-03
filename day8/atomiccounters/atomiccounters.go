package main
//The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

	var ops uint64
	//We’ll use an unsigned integer to represent our (always-positive) counter.

	var wg sync.WaitGroup
	//A WaitGroup will help us wait for all goroutines to finish their work.

    for i := 0; i < 50; i++ {
		//We’ll start 50 goroutines that each increment the counter exactly 1000 times.
        wg.Add(1)

        go func() {
			//To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()//Wait until all the goroutines are done

	fmt.Println("ops:", ops)
	//It’s safe to access ops now because we know no other goroutine is writing to it. Reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64.
}