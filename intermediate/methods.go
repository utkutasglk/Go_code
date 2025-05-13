package intermediate

import "fmt"

type Shape struct{
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with receiver value
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.length *= factor
}

func main() {

	rect := Rectangle{
		length: 5,
		width:  12,
	}
	area := rect.Area()
	fmt.Println(area)
	rect.Scale(2)
	area = rect.Area()

	fmt.Println(area)

	num := MyInt(-5)
	num1 := MyInt(9)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 10,width: 9}}
	fmt.Println(s.Area())
	

}

type MyInt int

// Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string{
	return "Welcome to MyInt Type"
}
