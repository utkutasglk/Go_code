package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader){

	buf := make([]byte,1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("error reading from reader", err)
	}
	fmt.Println(string(buf[:n]))

}

func writeToWriter(w io.Writer, data string){
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error reading from reader", err)
	}
}

func closeResource(c io.Closer,){
	err := c.Close()
	if err != nil {
		log.Fatalln("Error reading from reader: ", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // stack
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multipleReaderExample(){
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader("World")
	mr := io.MultiReader(r1,r2)
	buf := new(bytes.Buffer) // heap
	_,err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
	fmt.Println(buf.String())
}

func pipeExample(){
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filePath string, data string){
	file,err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		log.Fatalln("Error opening/creating file: ", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error opening/creating file: ", err)
	}


	// type(value)
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("error opening/closing file: ", err)
	// }
}




func main() {

	fmt.Println("==== Read from Reader ====")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("==== Write to Writer ====")
	var writer bytes.Buffer
	writeToWriter(&writer, "hello writer")
	fmt.Println(writer.String())

	fmt.Println("==== Buffer Example ====")
	bufferExample()

	fmt.Println("==== Multi Reader Example ====")
	bufferExample()

	fmt.Println("===== Pipe  Example =====")
	pipeExample()

	filePath := "io.txt"
	writeToFile(filePath,"Hello File!") 

	resource := &MyResource{name: "TestResource"}
	closeResource(resource)

}

type MyResource struct{
	name string
}

func (m MyResource) Close() error{
	fmt.Println("Closing resource: ", m.name)
	return nil
}