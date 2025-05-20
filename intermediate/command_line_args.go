package intermediate

import (
	"flag"
	"fmt"
	"os"
)


func main() {

	fmt.Println("Command", os.Args[0])

	for i,arg := range os.Args{
	fmt.Println("Argument1:",i,":",arg)
	}

	// define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name,"name","John","Name of the user")
	flag.IntVar(&age,"age",8,"Age of the user")
	flag.BoolVar(&male,"male",true,"Gender of user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)



}