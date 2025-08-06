package intermediate

import "fmt"

func pointers() {
	//var ptr *int

	var a = 10

	var ptr *int = &a

	fmt.Println(ptr)
	fmt.Println(&a)
	fmt.Println(*ptr)

	// nil pointer
	// var nilPtr *int
	modifyValue(ptr)

	fmt.Println(a)
	
}

func modifyValue(ptr *int) {
	*ptr++
}
