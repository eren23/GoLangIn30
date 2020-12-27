package main

import "fmt"

//

func main() {

    var a [5]int //  array of 5 ints, the type of the array and the length are both an array's type. By default an array is zero-valued, which for ints means 0s.
    fmt.Println("emp:", a)

    a[4] = 100 // set the 4th index
    fmt.Println("set:", a)
    fmt.Println("get:", a[4]) //get the 4th index

    fmt.Println("len:", len(a)) //length of the array

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int // Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}