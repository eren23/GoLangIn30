package main

import "fmt"

//Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

func main() {

	s := make([]string, 3)
	//Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
    fmt.Println("emp:", s)

    s[0] = "a" // We can set and get just like with arrays.
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

	s = append(s, "d") 
	// In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

	c := make([]string, len(s))
	//Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
    copy(c, s)
    fmt.Println("cpy:", c)

	l := s[2:5]
	// Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
    fmt.Println("sl1:", l)

	l = s[:5]
	//This slices up to (but excluding) s[5].
    fmt.Println("sl2:", l)

	l = s[2:]
	//And this slices up from (and including) s[2].
    fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	//We can declare and initialize a variable for slice in a single line as well.
    fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	//Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
	fmt.Println("2d: ", twoD)
	
	//Further read: https://blog.golang.org/slices-intro
}