package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1

	// keyword to filter lines
	keyword := "important"

	// read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line,keyword, "necessary")
			fmt.Printf("%d Filtered line: %v\n", lineNumber,line)
			lineNumber++
			fmt.Printf("%d Updated line: %v\n",lineNumber,updatedLine)
			lineNumber++
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("error scanning file:", err)
		return
	}
}
