package main

import (
	"log"
	// import the package helpers
	"packages/helpers"
)

func main() {
	log.Println("Hello")

	// Okay, what happens here
	//  I assign the variable myVar to be of the custom type struct SomeType
	// from the package helpers, which is imported
	var myVar helpers.SomeType
	// Now I set the TypeName inside the custom type SomeType
	// to some name
	myVar.TypeName = "Some name"
	// Now i can print out the Type name from the myVar variable
	log.Default().Println(myVar.TypeName)
}
