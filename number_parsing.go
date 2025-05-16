package main

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)

	if err != nil {
		fmt.Println("Error parsing the value :", err)
	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("Parsed Integer:", num+1)

	numStrInt, err := strconv.ParseInt(numStr,10,32)

	if err != nil {
		fmt.Println("Error parsing the value :", err)
	}
	fmt.Println("Parsed Integer: ", numStrInt)

	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr,64)
	if err != nil {
		fmt.Println("Error parsing the value :", err)
	}
	fmt.Printf("Parsed float: %.2f", floatVal)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr,2,64)
	if err != nil {
		fmt.Println("Error parsing binary value :", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr,16,64)
	if err != nil {
		fmt.Println("Error parsing binary value :", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", hex)

	invalidNum := "456aab"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing binary value :", err)
		return
	}
	fmt.Println("Parsed invalid number:", invalidParse)

}
