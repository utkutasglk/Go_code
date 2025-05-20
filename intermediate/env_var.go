package intermediate

import (
	"fmt"
	"os"
	"strings"
)


func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var:", user)
	fmt.Println("Home env var:", home)

	err := os.Setenv("FRUIT","APPLE")
	if err != nil {
		fmt.Println("Error setting environment variable: ", err)
	}
	fmt.Println("FRUIT:", os.Getenv("FRUIT"))

	for _, e := range os.Environ(){
		kvpair := strings.SplitN(e,"=",2)
		fmt.Println(kvpair[1])
	}

	os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable: ", err)
		return 
	}

	fmt.Println("FRUIT:", os.Getenv("FRUIT"))




}