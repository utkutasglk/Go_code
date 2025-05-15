package intermediate

import (
	"fmt"
	"regexp"
)


func main() {

	fmt.Println("He said, \"I am great\"")
	fmt.Println(`He said, "I am great"`)

	// Compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// test some strings
	email1  := "utku.tasguluk@gmail.com"
	email2 := "invalid_email"

	// match the strings 
	fmt.Println("Email1:", re.MatchString(email1))	
	fmt.Println("Email2:", re.MatchString(email2))

	// Capturing Groups
	// Compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d){4}-(\d{2})-(\d{2})`)

	// test string
	date :=" 2024-07-30"

	// find all submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// Source string
	str := "Hello World"

	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensetive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	// re = regexp.MustCompile(`go`)

	// Test string
	text := "Golang is great"

	// Match
	fmt.Println("Match:", re.MatchString(text))

}