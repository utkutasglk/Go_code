package basics

import "fmt"


func main() {


// q := func(x,y int)int{
// return x+y
// }

// operation := add

// result := operation(5,3)

// fmt.Println(result)



// fmt.Println(add(45,44),q(456,675))

// Passing a function as an argument
result := applyOperation(5,3, add)
fmt.Println(result)

// Returning and using a function
multiplyBy2 := createMultiplier(2)
fmt.Println("6*2 = ", multiplyBy2(6))


}

func add(a,b int) int{
return a+b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int,int)int)int{
	return operation(x,y)
}

// Function that returns a function
func createMultiplier(factor int) func(int)int{
	return func(x int) int{
		return x * factor
	}
}

