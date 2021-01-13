package main
// Throughout program execution, we often want to create data that isn’t needed after the program exits. Temporary files and directories are useful for this purpose since they don’t pollute the file system over time.
import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    f, err := ioutil.TempFile("", "sample")
	check(err)
	//The easiest way to create a temporary file is by calling ioutil.TempFile. It creates a file and opens it for reading and writing. We provide "" as the first argument, so ioutil.TempFile will create the file in the default location for our OS.

	fmt.Println("Temp file name:", f.Name())
	//Display the name of the temporary file. On Unix-based OSes the directory will likely be /tmp. The file name starts with the prefix given as the second argument to ioutil.TempFile and the rest is chosen automatically to ensure that concurrent calls will always create different file names.

	defer os.Remove(f.Name())
	//Clean up the file after we’re done. The OS is likely to clean up temporary files by itself after some time, but it’s good practice to do this explicitly.

    _, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	//We can write some data to the file.


    dname, err := ioutil.TempDir("", "sampledir")
    check(err)
	fmt.Println("Temp dir name:", dname)
	//If we intend to write many temporary files, we may prefer to create a temporary directory. ioutil.TempDir’s arguments are the same as TempFile’s, but it returns a directory name rather than an open file.

    defer os.RemoveAll(dname)

    fname := filepath.Join(dname, "file1")
    err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
	//Now we can synthesize temporary file names by prefixing them with our temporary directory.
}