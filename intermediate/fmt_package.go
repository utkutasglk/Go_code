package intermediate

import "fmt"


func main() {

	// Printing Functions

	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print(12,46)

	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(12,46)

	name := "James"
	age := 28

	fmt.Printf("Name: %s, Age: %d \n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// Formatting Functions
	s := fmt.Sprint("Hello", "World", 123,456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World", 123,456)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Print(sf)
	fmt.Println(sf)

	// Scanning Functions
	var name2 string
	var age2 int

	fmt.Println("What is your name ?")
	// fmt.Scan(&name2)
	fmt.Scanf("%s",name2)

	fmt.Println("What is your age ?")
	// fmt.Scan(&age2)
	fmt.Scanf("%d",age2)
	fmt.Printf("Name: %s, Age: %d\n", name2,age2)

	// error formatting functions
	err := checkAge(15)

	if err != nil {
		fmt.Println("error :", err)
	}

}

func checkAge(age int) error{
	if age < 18{
		return fmt.Errorf("age %d is too young to drive", age)
	}
	return nil
}