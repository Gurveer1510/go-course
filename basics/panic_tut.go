package basics

import "fmt"

func panic_tut() {
	process(20)
}

func process(input int) {

	defer fmt.Println("Deffered 1")
	defer fmt.Println("Deffered 2")
	if input < 0 {
		panic("Input must be a non negative number")
	}
	fmt.Println("Pocessing input: ", input)
	defer fmt.Println("Last deffered statement")
}
