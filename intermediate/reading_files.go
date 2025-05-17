package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file.")
		file.Close()
	}()

	fmt.Println("File opened successfully.")

	// // read the contents of the open file
	// data := make([]byte,1024)  // buffer to read data info
	// _,err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading data from file", err)
	// 	return
	// }

	// fmt.Println("file content : ", string(data))
	scanner := bufio.NewScanner(file)

	// read line by libe
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line:",line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:",err)
		return
	}

}
