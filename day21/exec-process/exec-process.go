package main
//In the previous example we looked at spawning external processes. We do this when we need an external process accessible to a running Go process. Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one. To do this we’ll use Go’s implementation of the classic exec function.
import (
    "os"
    "os/exec"
    "syscall"
)

func main() {

    binary, lookErr := exec.LookPath("ls") //For our example we’ll exec ls. Go requires an absolute path to the binary we want to execute, so we’ll use exec.LookPath to find it (probably /bin/ls).
    if lookErr != nil {
        panic(lookErr)
    }//For our example we’ll exec ls. Go requires an absolute path to the binary we want to execute, so we’ll use exec.LookPath to find it (probably /bin/ls).
	

    args := []string{"ls", "-a", "-l", "-h"} //Exec requires arguments in slice form (as apposed to one big string). We’ll give ls a few common arguments. Note that the first argument should be the program name.
    env := os.Environ()//Exec also needs a set of environment variables to use. Here we just provide our current environment.

    execErr := syscall.Exec(binary, args, env)//Here’s the actual syscall.Exec call. If this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process. If there is an error we’ll get a return value.
    if execErr != nil {
        panic(execErr)
    }
}