package basics

import (
	"fmt"
	"os"
)


func main() {

	defer fmt.Println("Defer statement")

	fmt.Println("Starting main function... ")

	// Exit with status code of 1
	os.Exit(1)

	// This will never be executed 
	fmt.Println("End of the main function")


}