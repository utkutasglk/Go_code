package main

import "fmt"


func main() {

	message := "Hello World"

	for i,v := range message{
		fmt.Printf("Index: %-3d, Rune: %-15c \n",i,v)
	}



}