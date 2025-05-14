package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


func main() {

	str := "Hello Go"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"

	result1 := str1 + " " + str2
	fmt.Println(result1)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// str conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// string splitting 
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits,",")
	parts2 := strings.Split(fruits1,"-")
	fmt.Println(parts)
	fmt.Println(parts2)

	countries := []string{"Germany","Turkey", "France"}
	joined := strings.Join(countries,", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str,"Go"))

	replaced := strings.Replace(str, "Go", "World", 1)
	fmt.Println(replaced)

	strwspace := "  Hello Everyone!   "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	fmt.Println(strings.Repeat("foo ",3))

	fmt.Println(strings.Count("hello","l"))
	fmt.Println(strings.HasPrefix("Hello","He"))
	fmt.Println(strings.HasSuffix("Hello","lo"))
	fmt.Println(strings.HasSuffix("Hello","la"))

	str5 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5,-1)
	fmt.Println(matches)

	// ----- STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World")

	// Convers this builder to string
	result := builder.String()
	fmt.Println(result)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you ?")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting Fresh")
	result = builder.String()
	fmt.Println(result)






}