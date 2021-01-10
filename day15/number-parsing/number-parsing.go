package main

import (
    "fmt"
    "strconv"
)

func main() {

    f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	//With ParseFloat, this 64 tells how many bits of precision to parse.

    i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	//For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	//ParseInt will recognize hex-formatted numbers.

    u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	//A ParseUint is also available.

    k, _ := strconv.Atoi("135")
	fmt.Println(k)
	//Atoi is a convenience function for basic base-10 int parsing.

    _, e := strconv.Atoi("wat")
	fmt.Println(e)
	//Parse functions return an error on bad input.
}