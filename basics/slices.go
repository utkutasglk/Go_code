package basics

import (
	"fmt"
)


func main() {

	// var numbers []int
	// var numbers1 = []int{1,2,3}

	// numbers2 := []int{9,8,7}

	// slice := make([]int,5)

	a := [5]int{1,2,3,4,5}
	slice1 := a[1:4]
	slice1 = append(slice1, 56)

	fmt.Println(slice1)

	for i,v := range a{
		fmt.Println(i,v)
	}

	fmt.Println(slice1)
	slice1[3] = 50

	// 



}