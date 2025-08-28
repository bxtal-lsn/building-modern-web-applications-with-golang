package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	names := []string{"one", "bxtal", "lsn"}
	log.Println(names)

	numbers := []int{1, 2, 5, 6, 10}

	log.Println(numbers)

	log.Println(numbers[0:2])

	// Go uses slices
	var myString string
	myString = "fish"
	log.Println(myString)
	// I want to store something that carries more than just a variable
	var mySlice []string

	mySlice = append(mySlice, "bxtal")
	mySlice = append(mySlice, "lsn")
	mySlice = append(mySlice, "mary")

	var MyIntSlice []int

	MyIntSlice = append(MyIntSlice, 2)
	MyIntSlice = append(MyIntSlice, 1)
	MyIntSlice = append(MyIntSlice, 3)

	sort.Ints(MyIntSlice)
	log.Println(mySlice)
	log.Println(MyIntSlice)

	// Map to Type struct
	myStructMap := make(map[string]User)

	me := User{
		FirstName: "bxtal",
		LastName:  "lsn",
	}

	myStructMap["me"] = me

	log.Println(myStructMap["me"].FirstName)

	// Map string int
	myIntMap := make(map[string]int)

	myIntMap["First"] = 1
	myIntMap["Second"] = 2

	log.Println(myIntMap["First"], myIntMap["Second"])

	// String map, are key value stores
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	//	var myString string
	//	var myInt int
	//
	//	myString = "Hoi"
	//	myInt = 11
	//
	//	mySecondString := "my second string"
	//
	//	log.Println(myString, myInt, mySecondString)

}
