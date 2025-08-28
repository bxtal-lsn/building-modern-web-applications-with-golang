package main

import "log"

func main() {
	// switch statement
	myVar2 := "cat"
	switch myVar2 {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}

	// number
	myNum := 100
	isTrueNum := true
	if myNum > 99 && isTrueNum {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrueNum {
		log.Println("1")
	} else if myNum == 101 || isTrueNum {
		log.Println("2")
	} else if myNum > 1000 && !isTrueNum {
		log.Println("3")
	}

	// cat
	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	// is true
	var isTrue bool

	isTrue = true

	if isTrue {
		log.Println("is true", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
}
