package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func StringsNRunes() {
	message := "Hello \n GO"
	rawMwssage := `Hello\nGO`

	fmt.Println(message)
	fmt.Println(rawMwssage)

	fmt.Println("Length of message: ", len(message))
	fmt.Println("The first character in message var is :", message[0])

	greeting := "Hello"
	name := "Gurveer"
	fmt.Println(greeting + name)

	str1 := "Apple"
	str2 := "Banana"

	fmt.Println(str1 < str2)

	// smallcase letter has greater value than uppercase letters

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c \n", i, char)
		fmt.Printf("%x\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	//strings are immutable
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)


	//Rune alias for int32

	// var ch rune = 'a'
	var jch rune = 26085

	fmt.Println(string(jch))
}
