package intermediate

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	Address
}

func (p Person) fullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) modifyAge() {
	p.age += 10
}

type Address struct {
	city string
	country string
}

func structs() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       21,
		Address: Address{
			city: "Delhi",
			country: "India",
		},
	}
	p1 := Person{}

	fmt.Println(p.firstName)
	fmt.Println(p.city)
	fmt.Println(p.country)
	fmt.Println(p.Address)
	fmt.Println(p1.age)

	// Anonymous struct

	// user := struct {
	// 	username string
	// 	email string
	// } {
	// 	username: "Gurveer",
	// 	email: "gurveer@yopmail.com",
	// }

	fmt.Println(p.fullName())
	p.modifyAge()
	fmt.Println(p.age)

}

