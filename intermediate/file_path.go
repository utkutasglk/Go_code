package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)


func main() {

	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	// join paths using filepath
	joinedPath := filepath.Join("home","Documents","downloads","file.zip")
	fmt.Println("Joined path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized path:", normalizedPath)

	dir,file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Path:", dir)
	fmt.Println(filepath.Base("/home/user/docs/"))

	fmt.Println("Is relativePath variable absolute: ",filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute:",filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file,filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Absolute path:", absPath)
	}
	
}