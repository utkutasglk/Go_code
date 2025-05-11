package main

import "fmt"


func main() {

	fmt.Println(factoriel(5))
	fmt.Println(factoriel(10))

	fmt.Println(sumOfdigits(9999))


}

func factoriel(n int) int{
	// Base case: factoriel of 0 is 1

	if n == 0{
		return 1
	}

	// Recursive case: factoriel of n is 
	return n * factoriel(n - 1)
}

func sumOfdigits(n int)int{
	// Base case
	if n < 10 {
		return n
	}

	return n%10 + sumOfdigits(n/10)

}