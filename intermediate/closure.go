package intermediate

import "fmt"

func main() {

	// sequence := adder()

	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
	// fmt.Println(sequence2())

	subtracter := func() func(int) int {

		countdown := 99
		return func (x int) int{
			countdown -= x
			return countdown
		}
	}()

	// Using closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))




}


func adder() func() int{

	i := 0 
	fmt.Println("Previous value of i :", i)
	return func() int{
	  i++
		fmt.Println("Added 1 to i :")
		return i
	}
	
}