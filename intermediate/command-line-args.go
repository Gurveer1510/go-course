package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func command_line_args() {
	fmt.Println("Command", os.Args[0])

	for i, arg := range os.Args {
		fmt.Println("Argument", i, ":", arg)
	}

	// Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "Gurveer", "Set the name of the user")
	flag.IntVar(&age, "age", 18, "Set the age of the user")
	flag.BoolVar(&male, "male", true, "Set the value for the male")

	flag.Parse()

	fmt.Println("Name", name)
	fmt.Println("Age", age)
	fmt.Println("Male", male)
}
