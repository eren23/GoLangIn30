package main

import "fmt"

//Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

func main() {

    m := make(map[string]int) // To create an empty map, use the builtin make: make(map[key-type]val-type).

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"] //Get a value for a key with name[key].
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m)) //The builtin len returns the number of key/value pairs when called on a map.

    delete(m, "k2") //The builtin delete removes key/value pairs from a map.
    fmt.Println("map:", m)

    _, prs := m["k2"]
	fmt.Println("prs:", prs)
	//The optional second return value when getting a value from a map indicates if the key was present in the map. 
	//This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.

    n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	//You can also declare and initialize a new map in the same line with this syntax.
}