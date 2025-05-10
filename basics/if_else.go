package basics

import "fmt"


func main() {

	// if condition {
	// 	// block of code 
	// }

	// age := 25

	// if age >= 18 {
	// 	fmt.Println("you are an adult")
	// }

	// tempareture := 25

	// if tempareture >= 30 {
	// 	fmt.Println("it is hot outside")
	// }else{
	// 	fmt.Println("It is cool outside")
	// }

	// score := 85

	// if score >= 90 {
	// 	fmt.Println("Grade A")
	// }else if score >= 80 {
	// 	fmt.Println("Grade B")
	// }else if score >= 70 {
	// 	fmt.Println("Grade C")
	// }else {
	// 	fmt.Println("Grade D")
	// }

	num := 18

	if num%2 == 0{
		if num%3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3.")
		}else{
			fmt.Println("Number is divisible by 2 but not 3")
		}
	}else{
		fmt.Println("Number is not divisible by 2.")
	}


}