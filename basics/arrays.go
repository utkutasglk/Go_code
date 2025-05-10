package basics

import "fmt"


func main() {

	var array [2]int 
	array[0]=1
	array[1]=2

	fmt.Println(array)

	array2 := [3]int{234,243,465}
	fmt.Println(array2)

	originalArray := []int{1,2,4,4,5,45,45,4,543,5,4,53,45,3645}

	for i := 0; i < len(originalArray); i++{

		fmt.Println(originalArray[i])
	}

	for _,v := range originalArray{
		fmt.Println(v)
	}

	a,_ := someFunction()

	fmt.Println(a)

	b := 2
	
	_ = b

	array3 := [3]int{1,2,3}
	array4 := [3]int{1,2,3}

	fmt.Println(array3==array4)



}

func someFunction()(int, int){
	return 1,2
}