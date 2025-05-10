package basics

import (
	"errors"
	"fmt"
)



func main() {

	q, r := divide(10,3)
	fmt.Printf("Quotient: %d. Remainder: %d\n", q,r)

	result, error := compare(4,4)

	if error != nil{
		fmt.Println("Error:", error)
	}else{
		fmt.Println(result)
	}
}




func divide(a,b int)(quotient int, remainder int){
	quotient = a/b
	remainder = a%b
	return 
}

func compare(a,b int)(string,error){
	if a > b{
		return "a is greater than b", nil
	} else if b > a{
		return "b is greater than a", nil
	} else{
		return "", errors.New("unable to compare which is greater")
	}
}