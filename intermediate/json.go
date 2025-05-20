package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email,omitempty"`
	Address Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	person := Person{
		Name: "utku",
		Age:  4,
	}
	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{
		Name:  "pangulot",
		Age:   55,
		Email: "pangulot@gmail.com",
		Address: Address{
			City:  "istanbul",
			State: "marmara",
		},
	}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
	}
	fmt.Println(string(jsonData1))

	jsonData2 := `{"full_name": "Harun Kaya", "emp_id":"0009","age":30,"address":{"city":"istanbul","state":"CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	fmt.Println(employeeFromJson)
	fmt.Println(employeeFromJson.Age)
	fmt.Println(employeeFromJson.Address.City)
	fmt.Println(employeeFromJson.FullName)

	listOfCityState := []Address{
		{City: "New York",State: "NY"},
		{City: "San Jose",State: "SJ"},
		{City: "Las Vegas",State: "LA"},
		{City: "Modeste",State: "CA"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshalling to Json", err)

	}
	fmt.Println(string(jsonList))

	// Handling unkown json structures
	jsonData3 := `{"name":"Ufuk", "age":30,"address":{"city":"New York", "state":"NY"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData3),&data)
	if err != nil {
		log.Fatalln("Error Unmarshalling JSON:", err)
	}
	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])



}

type Employee struct{
	FullName string `json:"full_name"`
	EmID string `json:"emp_id"`
	Age int `json:"age"`
	Address Address `json:"address"`
}
