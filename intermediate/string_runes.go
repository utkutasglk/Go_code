package intermediate

import (
	"fmt"
	"unicode/utf8"
)


func main() {

	message := "Hello, \nGo"
	message1 := "Hello, \tGo"
	message2 := "Hello, \rGo"


	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println(len(message))

	fmt.Println(message[0])

	greeting := "Hello "
	name := "Utku"
	fmt.Println(greeting+name)

	str1 := "Apple"
	str2 := "banana"
	fmt.Println(str1 < str2)

	for i, char := range message{
		fmt.Printf("Character at index %d is %c\n", i,char)
	}

	fmt.Println("rune :", utf8.RuneCountInString(message))

	greetingWithName := greeting+name
	fmt.Println(greetingWithName)

	var ch rune = 'c'
	jch := '上'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n",ch)
	fmt.Printf("%c\n",jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n",cstr)

	const NIHONGO = "宛よぇ"
	fmt.Println(NIHONGO)

	jhello := "宛よぇ"

	for _,runeValue := range jhello{
		fmt.Printf("%c\n",runeValue)
	}





}