package basics

import "fmt"

var middlename = "Singh"

func variables () {
	// var age int
	// var name string = "Gurveer"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Singh"
	middlename = "narr"
	// Default values
	// Numeric Types: 0
	// Boolen Types: false
	// String Type: ""
	//Pointers, slices, maps, function, and struct : nil

	// .... SCOPE

	fmt.Println(middlename)

}

func printName() {
	firstName := "Gurveer"
	fmt.Println(firstName)
}