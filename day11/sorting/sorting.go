package main

import (
    "fmt"
    "sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	//Sort methods are specific to the builtin type; here’s an example for strings. Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	//An example of sorting ints.
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	//We can also use sort to check if a slice is already in sorted order.
    fmt.Println("Sorted: ", s)
}