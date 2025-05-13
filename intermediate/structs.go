package intermediate

import "fmt"

type Person struct{

	firstName string
	lastName string 
	age int

}



func main() {

	
	utku := Person{
		firstName: "Utku",
		lastName: "Tasguluk",
		age: 27,
	}

	harun := Person{
		firstName: "Harun",
		lastName: "Kaya",
		age: 18,
	}


	type Car struct{
		Brand string
		Price int
	}


	sahin := Car{
		Brand:"Tofas",
		Price: 25000,
	}

	fmt.Println(utku.firstName)
	fmt.Println(sahin.Brand)
	fmt.Println(harun.lastName)

	
	// Anonymous struct
	user := struct{
		userName string
		email string
	}{
		userName: "harun123",
		email: "harunkaya@gmail.com",
	}

	fmt.Println(user)

	fmt.Println(utku.fullName())

	harun.gettingOld()

}

// method 1 
func (a Person) fullName() string{
	return a.firstName +" "+ a.lastName
}


// method 2 
func (a Person) gettingOld() {
	fmt.Println("My age will be ", (a.age)," years old")
}

// method3 
func (a *Person) increaseAgeBy1(){
	a.age++
}

