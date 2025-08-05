package basics

import "fmt"

func recover_tut() {
	process1()
	fmt.Println("Return from process")
}

func process1() {
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("REcovered: ", r)
		}
	}()

	fmt.Println("Start PRocess")
	panic("Something went wrong")
	fmt.Println("End Process")
}
