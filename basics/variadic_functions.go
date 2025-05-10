package basics

import "fmt"


func main() {

	// fmt.Println("sum of 1,2,3 = ", sum(1,2,3,))

	returnS, result := sum("The sum of number 1,2,3 :", 1,2,3)
	fmt.Println(returnS,result)

	numbers := []int{1,2,3,4,5,6,7,}
	sequence, total := sum("utku",numbers...)
	fmt.Println(sequence,total)


}

func sum(returnString string,nums ...int)(string,int){
	total := 0

	for _, v := range nums{
		total += v
	}
	return returnString,total

}

