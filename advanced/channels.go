package advanced

import (
	"fmt"
	"time"
)

func main() {

	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "World"

		for _,e := range "abcde"{
			greeting <- "Alphabet" + string(e)
		}
	}()

	// go func ()  {
	// 	receiver := <- greeting
	//   fmt.Println(receiver)
	// 	receiver = <- greeting
	//   fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <- greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program.")

}
