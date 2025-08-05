package basics

import (
	"fmt"
	"os"
)

func exit_tut() {
	defer fmt.Println("Deffered statement")
	fmt.Println("Starting the main function ")

	//exit  with status code of 1

	os.Exit(1)

	// this will never be executed

	fmt.Println("End of main function")
}
