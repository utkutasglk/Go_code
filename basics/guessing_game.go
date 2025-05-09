package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	source := rand.NewSource(time.Now().UnixNano())

	random := rand.New(source)

	// Generate random and random number 
	target := random.Intn(100) + 1

	// welcome message
	fmt.Println("Welcome to the guessing game ")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is ?")

	var guess int 

	for{
		fmt.Println("Enter your guess : ")
		fmt.Scanln(&guess)

		// Check if the guess is correct
		if guess == target {
			fmt.Println("you guess is correct")
			break
		}else if guess < target {
			fmt.Println("guess is to low ! Try guessing a higher number.")
		}else {
			fmt.Println("guess is to high ! Try guessing a lower number.")
		}


	}


}