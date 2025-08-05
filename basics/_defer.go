package basics

import "fmt"

func deffer() {
	process(10)
}

func process(i int) {
	defer fmt.Println("Deffered i value: ", i)
	i++
	defer fmt.Println("1st Defered statement executed")
	defer fmt.Println("2nd Defered statement executed")
	fmt.Println("Normal execution")
	fmt.Println("Actual value: ", i)
}
