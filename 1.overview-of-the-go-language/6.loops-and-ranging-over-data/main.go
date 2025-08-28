package main

import "log"

func main() {
	// loop over a custom type
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	// set variable users as type slice of User
	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Dow", "mary@dow.com", 31})
	users = append(users, User{"John", "Snow", "john@snow.com", 32})
	users = append(users, User{"test", "Testsen", "test@testsen.com", 30})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
	// loop over string
	// NOTE a string is a slice of bytes!
	// A Rume is a Byte
	// strings are actually immutable, so if you assign a new value to a string, the previous variable is destroyed
	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	// Range over a map
	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "MoccaFido"

	for animalType, moochie := range animals2 {
		log.Println(animalType, moochie)
	}

	// a slice of strings, unsure how much data is there
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// The first value is the current iteration
	// The second one is the value of what i am ranging over
	for i, animal := range animals {
		log.Println(i, animal)
	}

	// If i do not care about the iteration number:
	// I can use the underscore instead, it is a blank identifier
	for _, animal := range animals {
		log.Println(animal)
	}

	// basic loop
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}
}
