package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// registers the handle function
	// First argument is the patterm. Here it is set to home
	// The second argument requires a handler function
	// which is the response/request pattern
	// First we specify the response writer as w, and a pointer to http.request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// format Fprintf takes two arguments
		// first a response writer, in this case w
		// and then a string, in this case hello world
		// the function can then recieve any number of arguments after that

		// The function returns a number of bytes, and an error
		// Number of bytes can be ignored with _
		n, err := fmt.Fprintf(w, "Hello World!\n")
		// Error handling
		if err != nil {
			// if error, print it
			fmt.Println(err)
		}
		log.Printf("Number of bytes written: %d\n", n)
	})
	// Start http server
	// specify port, and if golang mux is used, then nil as hendle func
	// also, ignore the error, as the program will shut down if error
	_ = http.ListenAndServe(":8080", nil)
}
