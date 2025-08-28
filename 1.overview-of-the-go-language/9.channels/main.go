package main

import (
	"log"

	"github.com/bxtal-lsn/channels/helpers"
)

const numPool = 1000

// this func takes an argument of intChan, and is of type chan of int
func CalculateValue(intChan chan int) {
	// i want to generate a random number
	// and pass it to a calculator

	// get random number from random number generator
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
func main() {
	// Channels

	// Create a channel, a place that can send information
	// to be recieved in 1 or more places in my go program
	// this channel can only hold ints
	intChan := make(chan int)

	// I want to close this when the function is finished
	// defer is whatever comes after defer, execute as soon as the current function is done
	// db conn should not be open forever. limited number of db conns at a given time
	// important to close conns, always do this
	defer close(intChan)

	// start a go routine, why enables concurrent processes
	// use go in front of the function you want to call
	go CalculateValue(intChan)

	// Now i need to listen to the response of that channel
	num := <-intChan

	log.Println(num)
}
