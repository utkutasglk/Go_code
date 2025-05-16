package intermediate

import (
	"fmt"
	"math/rand"
)

func main() {

	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(rand.Intn(6) + 5)
	// fmt.Println(val.Intn(6) + 5)
	// fmt.Println(rand.Float64()) // between 0.0 and 1.0

	for {
		// Show the menu 
		fmt.Println("Welcome to the Dice game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Println("Enter your choice (1 or 2):")

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2:")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// show the results 
		fmt.Printf("you rolled a %d and a %d \n", die1, die2)
		fmt.Println("total:",die1 + die2)

		// ask the user wants to roll again
		fmt.Println("Do u want to roll again ? (y/n): ")
		var rollAgain string

		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}

		if rollAgain == "n" {
			fmt.Println("Thanks for playing ! goodbye")
			break
		}



	}
}
