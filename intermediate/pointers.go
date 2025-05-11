package intermediate

import "fmt"


func main() {

	// Pointer is a variable that store memory address of another variable
	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)  // dereferencing a pointer


}