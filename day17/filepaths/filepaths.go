package main
//The filepath package provides functions to parse and construct file paths in a way that is portable between operating systems; dir/file on Linux vs. dir\file on Windows, for example.

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {

    p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	//Join should be used to construct paths in a portable way. It takes any number of arguments and constructs a hierarchical path from them.

    fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	//You should always use Join instead of concatenating /s or \s manually. In addition to providing portability, Join will also normalize paths by removing superfluous separators and directory changes.

    fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	//Dir and Base can be used to split a path to the directory and the file. Alternatively, Split will return both in the same call.

    fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	//We can check whether a path is absolute.


    filename := "config.json"

    ext := filepath.Ext(filename)
	fmt.Println(ext)
	//Some file names have extensions following a dot. We can split the extension out of such names with Ext.

	fmt.Println(strings.TrimSuffix(filename, ext))
	//To find the file’s name with the extension removed, use strings.TrimSuffix.

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	//Rel finds a relative path between a base and a target. It returns an error if the target cannot be made relative to base.
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}