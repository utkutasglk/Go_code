package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)


func main() {

	// tmpl := template.New("example")

	// tmpl,err := template.New("example").Parse("Welcome, {{.name}}! How are you doing ?\n")
	// if err != nil {
	// 	panic(err)
	// }

	// tmpl := template.Must(template.New("Example").Parse("Welcome, {{.name}}! How are you doing ?\n"))


	// Define data for the welcome message template
	// data := map[string]interface{}{
	// 	"name":"john",
	// }

	// err := tmpl.Execute(os.Stdout, data)

	// if err != nil { 
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name :")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define named templates for different types of 
	templates := map[string]string {
		"welcome": "Welcome {{.name}}! We're glad you joined.", 
		"notification":"{{.nm}}, you have a new notification: {{.ntf}}",
		"error":"Oops! An error occured: {{.em}}",
	}

	// Parse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// show menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option:")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name":name}
		case "2":
			fmt.Println("Enter your notification message:")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"nm":name,"ntf": notification}
		
		case "3":
			fmt.Println("Enter your error message:")
			errorMessage,_ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"nm":name, "em": errorMessage}
		case "4":
			fmt.Println("Exiting ....")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			continue


		}

		// render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template: ", err)
		}
	}

}