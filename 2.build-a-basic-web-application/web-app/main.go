package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler fuck
// Another line
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is: %d", sum)
}

func divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.00, 10.0)
	if err != nil {
		fmt.Fprint(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by o")
		return 0, err
	}
	result := x / y
	return result, nil
}

// addValues add two numbers together and return the result
func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", divide)

	fmt.Printf("Starting application on port number %s\n", portNumber[1:])
	_ = http.ListenAndServe(portNumber, nil)
}
