package intermediate

import (
	"fmt"
	"math"
)

// Simple interface definition
type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rectangle1 struct {
	Width  float64
	Height float64
}

func (r Rectangle1) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle1) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())

}

func main() {

	rec := Rectangle{Width: 3, Height: 4}
	cir := Circle{Radius: 5}
	rec1 := Rectangle1{Width: 3, Height: 4}

	measure(rec)
	measure(cir)
	measure(rec1)

	MyPrinter(1, "john", 45.9, true)
	value("utku")
	value(4)

}

func value(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type:int")
	case string:
		fmt.Println("Type:string")
	default:
		fmt.Println("Type:Unknown")
	}
}

func MyPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
