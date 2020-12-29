package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful\n")
    name := r.FormValue("name")
    address := r.FormValue("address")

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // that one turns string path into a http.Dir type because Handle function expects smth like that

	http.HandleFunc("/hello", helloHandler) 
	http.HandleFunc("/form", formHandler)
	http.Handle("/", fileServer)
	

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "Hello!")
    // })

	fmt.Printf("Starting server at port 8080\n")
	if err:= http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	}
}