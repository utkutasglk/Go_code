package intermediate

import "fmt"

func main() {

	var a int = 32
	b := int32(a)
	c := float64(b)
	d := bool(true)

	fmt.Println(d)
	fmt.Println(c)

	// type(value)

	g := "Hello"
	var h []byte
	h = []byte(g)
	fmt.Println(h)

}
