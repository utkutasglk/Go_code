package basics

import "fmt"


func main() {

	// panic(interface{})

	// example of a valid input
	process(10)

	// example of invalid input
	process(-5)


}

func process(input int){

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0{
		fmt.Println("Before Panic")
		panic("input must be a non negative number")

	}
	fmt.Println("processing input:", input)
}