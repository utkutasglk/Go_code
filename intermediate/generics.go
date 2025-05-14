package intermediate

import "fmt"


func swap[T any](a,b T)(T,T){
	return b,a
}

type Stack[T any] struct{
	elements []T
}

func (s *Stack[T]) push(element T){
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop()(T,bool){

	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements) - 1]
	s.elements = s.elements[:len(s.elements) -1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool{
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Println("Stack elements : ")

	for _, v := range s.elements {
		fmt.Println(v)
	}

}

func main() {

	x, y := 1,2
	x, y = swap(x,y)
	fmt.Println(x,y)

	x1, y1 := "tasguluk","utku"
	x1, y1 = swap(x1,y1)
	fmt.Println(x1,y1)

	inStack := Stack[int]{}
	inStack.push(1)
	inStack.push(2)
	inStack.push(3)
	inStack.printAll()
	fmt.Println(inStack.pop())
	inStack.printAll()
	inStack.pop()
	inStack.pop()
	fmt.Println(inStack.isEmpty())


}