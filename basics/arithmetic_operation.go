package basics

import "fmt"

var age int = 45


func main(){

	var a,b int = 10,3
	var result int
	var c string = "utku"
	fmt.Println(c)

	fmt.Println(age)


	

	result = a + b
	fmt.Println("Addition:",result)

	result = a - b
	fmt.Println("Subtraction:",result)

	result = a * b
	fmt.Println("Multipication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22.0 / 7.0
	fmt.Println(p)

	var maxInt int64 = 9347865309845768
	fmt.Println(maxInt)



}