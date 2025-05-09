package main

import "fmt"


func main() {

	var a map[int]string

	b := map[int]int{
		4:5,
		6:7,
		8:9,
	}

	c := make(map[string]int)

	c["utku"] = 101
	c["ufuk"] = 202
	c["figen"] = 303

	// delete(c, "utku")
	// clear(c)

	fmt.Println(a,b,c)

	myMap := map[string]int{"a":1,"b":2,"c":3}
	fmt.Println(myMap)

	for i,v := range c{
		fmt.Println(i,v)
	}

}