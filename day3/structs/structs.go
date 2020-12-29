package main

import "fmt"

type person struct { //Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
    name string
    age  int
}

func newPerson(name string) *person { //newPerson constructs a new person struct with the given name

    p := person{name: name}
    p.age = 42 //You can safely return a pointer to local variable as a local variable will survive the scope of the function.
    return &p
}

func main() {

    fmt.Println(person{"Bob", 20}) //syntax for a new struct 

    fmt.Println(person{name: "Alice", age: 30}) // you also can name the fields while initilazing a struct

    fmt.Println(person{name: "Fred"}) // ommited fields will bpe zero valued

    fmt.Println(&person{name: "Ann", age: 40}) // & prefix yields a pointer to struct

    fmt.Println(newPerson("Jon")) // It’s idiomatic to encapsulate new struct creation in constructor functions

    s := person{name: "Sean", age: 50} // Access struct fields with a dot.
    fmt.Println(s.name)

    sp := &s			// You can also use dots with struct pointers - the pointers are automatically dereferenced.
    fmt.Println(sp.age)

    sp.age = 51 // Structs are mutable.
    fmt.Println(s.age)
}