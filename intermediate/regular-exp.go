package intermediate

import (
	"fmt"
	"regexp"
)

func regex_() {
	// Compile a regex to match an email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	email1 := "user@email.com"
	email2 := "invalid email"

	fmt.Println("Email 1:", re.MatchString(email1))
	fmt.Println("Email 2:", re.MatchString(email2))

	//Capturing Groups
	// Compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d){4}-(\d{2}-(\d{2}))`)

	// Test string
	date := "2024-07-30"

	// Find all matches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	str := "Hello world"

	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)


}
