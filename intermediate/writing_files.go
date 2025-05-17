package intermediate

import (
	"fmt"
	"os"
)


func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file", file)
		return
	}
	defer file.Close()

	// write data to the file
	data := []byte("Hello World\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
	fmt.Println("Data has been written to file successfully ")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go\n")
	if err != nil {
		fmt.Println("Error writing to file: ",err)
		return
	}


}