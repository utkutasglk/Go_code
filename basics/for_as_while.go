package basics

import "fmt"


func main() {

	// sum := 0

	// for{
	// 	sum += 10

	// 	if sum > 100 {
	// 		break
	// 	}
	// }
	// fmt.Println(sum)

	num := 0 

	for num <= 10{
		if num%2 == 0{
			num++
			continue
		}
		fmt.Println("Odd number:", num)
		num++
	}





}