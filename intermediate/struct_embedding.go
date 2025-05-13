package intermediate

import "fmt"

type Person struct{
	name string
	age int 
}

type Employee struct{
	Person // Embedded struct
	emId string
	salary float64
}

func (p Person) introduce() {
	fmt.Printf("Hi I'm %s and I'm %d years old.\n", p.name, p.age)
}


func (e Employee) introduce(){
	fmt.Printf("Hi I'm %s and I'%sm %.2f years old.\n", e.name, e.emId, e.salary)

}

func main() {

	emp := Employee{
		Person: Person{name: "utku",age: 45},
		emId: "E001",salary: 5000,
	}

	fmt.Println("Name",emp.name)
	fmt.Println("Age",emp.age)
	fmt.Println("EmId",emp.emId)
	fmt.Println("Salary",emp.salary)

	emp.introduce()


}