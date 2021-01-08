package main

import (
	//Go offers excellent support for string formatting in the printf tradition. Here are some examples of common string formatting tasks.
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
	fmt.Printf("%v\n", p)
	//Go offers several printing “verbs” designed to format general Go values. For example, this prints an instance of our point struct.

	fmt.Printf("%+v\n", p)
	//If the value is a struct, the %+v variant will include the struct’s field names.

	fmt.Printf("%#v\n", p)
	//The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.

	fmt.Printf("%T\n", p)
	//To print the type of a value, use %T.

	fmt.Printf("%t\n", true)
	//Formatting booleans is straight-forward.


	fmt.Printf("%d\n", 123)
	//There are many options for formatting integers. Use %d for standard, base-10 formatting.


	fmt.Printf("%b\n", 14)
	//This prints a binary representation.


	fmt.Printf("%c\n", 33)
	//This prints the character corresponding to the given integer.


	fmt.Printf("%x\n", 456)
	//%x provides hex encoding.


	fmt.Printf("%f\n", 78.9)
	//There are also several formatting options for floats. For basic decimal formatting use %f.


    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)//%e and %E format the float in (slightly different versions of) scientific notation.



    fmt.Printf("%s\n", "\"string\"")
	//For basic string printing use %s.



	fmt.Printf("%q\n", "\"string\"")
	//To double-quote strings as in Go source, use %q.



	fmt.Printf("%x\n", "hex this")
	//As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.


	fmt.Printf("%p\n", &p)
	//To print a representation of a pointer, use %p.


	fmt.Printf("|%6d|%6d|\n", 12, 345)
	//When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.


	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	//You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.



	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	//To left-justify, use the - flag.



	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	//You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.



	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	//To left-justify use the - flag as with numbers.



    s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	//So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.



	fmt.Fprintf(os.Stderr, "an %s\n", "error")
	//You can format+print to io.Writers other than os.Stdout using Fprintf.


}