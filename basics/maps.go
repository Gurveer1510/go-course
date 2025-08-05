package basics

import (
	"fmt"
	"maps"
)

func mapss() {

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["Gurveer"] = 21
	myMap["Karan"] = 14

	fmt.Println(myMap)

	delete(myMap, "Gurveer")

	fmt.Println(myMap)

	value, ok := myMap["Gurveer"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key not found")
	}

	myMap2 := map[string]int{"a": 1, "b": 2}
	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
	
	fmt.Println("Key \t\t Value")
	for k, v := range myMap3 {

		fmt.Printf("%v \t\t %v\n", k, v)
	}

}
