package main
//Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.
import (
    "flag"
    "fmt"
)

func main() {

	wordPtr := flag.String("word", "foo", "a string")
	//Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag word with a default value "foo" and a short description. This flag.String function returns a string pointer (not a string value); we’ll see how to use this pointer below.

    numbPtr := flag.Int("numb", 42, "an int")//This declares numb and fork flags, using a similar approach to the word flag.
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string //It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse() //Once all flags are declared, call flag.Parse() to execute the command-line parsing.

    fmt.Println("word:", *wordPtr) //Here we’ll just dump out the parsed options and any trailing positional arguments. Note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}