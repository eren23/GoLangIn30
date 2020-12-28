package main

import "fmt"

func vals() (int, int) {
	return 3, 7
	//The (int, int) in this function signature shows that the function returns 2 ints.
}

func main() {

    a, b := vals()
    fmt.Println(a)
	fmt.Println(b)
	//Here we use the 2 different return values from the call with multiple assignment.

    _, c := vals()
	fmt.Println(c)
	//If you only want a subset of the returned values, use the blank identifier _.
}