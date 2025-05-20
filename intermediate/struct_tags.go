package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
}

func main() {

	person := Person{FirstName: "Utku", LastName: "Tasguluk", Age: 28}

	json1, err := json.Marshal(person)

	if err != nil {
		log.Fatalln("Error is occured in json.", err)
	}

	fmt.Println(string(json1))

}
