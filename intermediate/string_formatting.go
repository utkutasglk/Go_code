package intermediate

import "fmt"

func main() {
	
	num := 425
	fmt.Printf("%05d\n", num)

	message := "Hello lkjsdhfgl"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \nWorld"
	message2 := `Hello \nWorld`

	fmt.Println(message1)
	fmt.Println(message2)
	
}