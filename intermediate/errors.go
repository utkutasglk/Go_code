package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// compute the square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty data")
	}

	// Process data
	return nil
}

func main() {

	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	result1, err1 := sqrt(-1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println(result1)

	data := []byte{}

	if err := process(data); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Data Processed Successfully")

	err2 := eprocess( )

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")


}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s ", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readData() error{
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error{
	return errors.New("config Error")

}