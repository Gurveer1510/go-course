package intermediate

import (
	"fmt"
	"strings"
)

func string_function() {
	// str := "Hello Go!"

	// fmt.Println(len(str))

	// str1 := "Hello"
	// str2 := "World"
	// result := str1 + " " + str2
	// fmt.Println(result)

	// fmt.Println(string(str[1]))
	// fmt.Println(str[1:7])

	// // String conversion
	// num := 19
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))

	// // string splitting
	// fruits := "mango,apple,orange"
	// parts := strings.Split(fruits, ",")
	// fmt.Println(parts)

	// // countries := make([]string, 1)
	// // countries = append(countries, "Germay", "France" , "India")
	// countries := []string{"India","Argentina","Scotland"}
	// joined := strings.Join(countries, ",")
	// fmt.Println(joined)

	// fmt.Println(strings.Contains(str, "Hel"))

	// replaced := strings.Replace(str, "Go", "World", 3)
	// fmt.Println(replaced)

	// strwspace := "Hello Everywone"
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace))

	// fmt.Println(strings.ToLower(strwspace))
	// fmt.Println(strings.ToUpper(strwspace))

	// fmt.Println(strings.Repeat("Hello", 1))

	// str5 := "Hello, 123 Go! 11"
	// re := regexp.MustCompile(`\d+`)
	// res := re.FindAllString(str5, 4)
	// fmt.Println(res)

	// str6 := "Hello Shorty"
	// fmt.Println(utf8.RuneCountInString(str6))

	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// convert builder to string
	result := builder.String()
	fmt.Println(result)

	// using writeRune to add a character

	builder.WriteRune(' ')
	builder.WriteString(" how are you  ?")
	result = builder.String()
	fmt.Println(result)

	// reset the builder
	builder.Reset()
	builder.WriteString("Fresh start!")
	result = builder.String()
	fmt.Println(result)
}
