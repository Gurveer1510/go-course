package basics

import "fmt"

func arrays() {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 11
	numbers[4] = 44
	fmt.Println(numbers)

	fruits := []string{"apple", "mange", "peach"}

	for _, v := range fruits {
		fmt.Println(v)
	}

	originalArray := [3]int{1,2,3}
	copiedArray := &originalArray

	copiedArray[0] = 100

	fmt.Println("original ", originalArray)
	fmt.Println("copied ", *copiedArray)
	
}
