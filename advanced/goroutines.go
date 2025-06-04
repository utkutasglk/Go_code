package advanced

import (
	"fmt"
	"time"
)


func main() {

	var err error

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func(){
	err = doWork()
	}()
	// err = go doWork() // This is not accepted
	go printNumbers()
	go printLetters()
	time.Sleep(2*time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Work completed successfully")
	}

}

func sayHello() {
	time.Sleep(1*time.Second)
	fmt.Println("Hello from goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:",i,time.Now())
		time.Sleep(100*time.Millisecond)
	}
}

func printLetters(){
	for _,letter := range "abcd"{
		fmt.Println(string(letter), time.Now())
		time.Sleep(200*time.Microsecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1*time.Second)

	return fmt.Errorf("an error occured in doWork")
}