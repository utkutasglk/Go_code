package main

import (
	"fmt"
)


func main() {

	// switch expression{
	// case value1:
	// 	// Code to be executed if expression
	// case value2:
	// 	// Code to be executed if expression
	// case value3:
	// 	// Code to be executed if expression
	// case value4:
	// 	// Code to be executed if expression
	// default:
		// 
	// }

	fruit := "pinapple"

	switch fruit {
	case "apple":
		fmt.Println("It is an apple.")
	case "banana":
		fmt.Println("It is a banana")
	default:
		fmt.Println("Unknown fruit")
	}

	//Multiple Conditions
	day := "monday"

	switch day {
	case "monday", "tuesday", "wednesday", "thursday","friday":
		fmt.Println("It is a weekday")
	case "sunday":
		fmt.Println("It is a weekend")
	default:
		fmt.Println("Invalid day.")
	}

	number := 15 

	switch{
	case number < 10:
		fmt.Println("Number is lower than 10")
		fallthrough
	case number > 10 && number <= 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	checkType(10)
	checkType(true)


}

func checkType(x interface{}){

	switch x.(type){
	case int:
		fmt.Println("It is an integer")
	case float32:
		fmt.Println("It is a float")
	case string:
		fmt.Println("it is a string")
	default:
		fmt.Println("Unknown Type")
	}
	
}