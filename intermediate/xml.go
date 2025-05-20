package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"-"`
	Address Address  `xml:"address"`
}

type Address struct{
	City    string   `xml:"city"`
	State   string   `xml:"state"`

}

func main() {

	person := Person{Name: "Utku", Age:30, Email: "utku.tasguluk@gmail.com", Address: Address{City: "istanbul", State: "Marmara"}}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling data into XML", err)
	}
	fmt.Println(string(xmlData))

	xmlData1, err := xml.MarshalIndent(person,""," ")
	if err != nil {
		log.Fatalln("Error marshalling data into XML", err)
	}
	fmt.Println(" XML Data with Indent\n",string(xmlData1))

	//xmlRaw := `<person><name>Harun</name><age>34</age></person>`
	xmlRaw := `<person><name>Harun</name><age>34</age><address><city>Konya</city><state>anadolu</state></address></person>`

	var personXml Person
	err = xml.Unmarshal([]byte(xmlRaw),&personXml)
	if err != nil {
		log.Fatalln("error unmarshalling xml:", err)
	}

	fmt.Println(personXml)
	fmt.Println("Local: ",personXml.XMLName.Local)
	fmt.Println("Space: ",personXml.XMLName.Space)

}
