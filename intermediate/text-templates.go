package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func text_templates() {
	// // tmpl := template.New("Example")
	// // tmpl, err := template.New("Example").Parse("Welcome, {{.name}}!, How are you doing ?")
	// tmpl := template.Must(template.New("Example").Parse("Welcome, {{.name}}!, How are you doing ?"))

	// //DEfine data fo rthe welcome message template
	// data := map[string]interface{}{
	// 	"name": "John",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter your name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"Welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "{{.name}} You have a new notification: {{.notification}}",
		"error":        "Oops! An error occurred: {{.error}}",
	}

	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show th menu
		fmt.Println("\nMenu")
		fmt.Println("\n1. Join")
		fmt.Println("\n2. Get Notification")
		fmt.Println("\nM3. Get Error")
		fmt.Println("\n4. Exit")
		fmt.Println("Choose an option")

		choice , _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1" :
			tmpl = parsedTemplates["Welcome"]
			data = map[string]interface{}{
				"name" : name,
			}
			err := tmpl.Execute(os.Stdout, data)
			if err != nil {
				fmt.Println("Error executing the template: ", err)
			}
		case "2":
			fmt.Println("Enter your notif message")
			notification , _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{
				"name" : name, "notification" : notification,
			}
			err := tmpl.Execute(os.Stdout, data)
			if err != nil {
				fmt.Println("Error executing the template: ", err)
			}
		case "3":
			fmt.Println("Enter your error message")
			errormsg , _ := reader.ReadString('\n')
			errormsg = strings.TrimSpace(errormsg)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{
				"name" : name, "error" : errormsg,
			}
			err := tmpl.Execute(os.Stdout, data)
			if err != nil {
				fmt.Println("Error executing the template: ", err)
			}
		case "4":
			return
		}
	}
}
