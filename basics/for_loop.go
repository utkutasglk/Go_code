package basics

import "fmt"


func main() {

	// Simple iteration over a range
	// for i := 0; i < 5; i++ {

	// 	fmt.Println("utku")
	// }

	// // iterate over collection 
	// numbers := []int{1,2,3,4,5,6}

	// for index,value := range numbers {
	// 	fmt.Printf("Index: %d, Value:%d\n\b", index,value)
	// }

	// fmt.Printf("Binary: %b\\%b", 40, 50) // Prints `Binary: 100\101`

	// for i := 1; i < 10; i++{

	// 	if i%2 == 0 {
	// 		continue // continu the loop but skip the rest of lines
	// 	}
	// 	fmt.Println("Odd number:", i)

	// 	if i==5{
	// 		break
	// 	}

	// }

	// rows := 5

	// // Outer loop
	// for i := 1; i<=rows; i++{
	// 	// inner loop for spaces before starts 
	// 	for j:=1; j <=rows-i; j++{
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for stars 
	// 	for k := 1; k <=2*i-1; k++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // Move to the next line

	// }

	s := "utku tasguluk"

	for i := range s {
		fmt.Println(i)
	}




}