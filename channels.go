package main

import (
	"fmt"
	"time"
)

func main() {
	// variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking cuz it's continuosly trying to recieve values
		// greeting <- "World!"
		for _, ele := range "abcde" {
			greeting <- "Alphabet: " + string(ele)
		}
	}()

	// go func () {
	// 	receiver := <- greeting
	// 	fmt.Println(receiver)
	// }()

	for range 5 {
		rcvr := <- greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program.")
}
