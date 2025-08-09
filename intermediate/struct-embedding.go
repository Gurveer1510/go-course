package intermediate

import "fmt"

type Person1 struct {
	name string
	age  int
}

type Employee struct {
	Person1
	empId  string
	salary float64
}

func (p Person1) introduce() {
	fmt.Printf("Hi I'm %s and I'm %d years old", p.name, p.age)
}

func (emp Employee) introduce() {
	fmt.Printf("Hi I'm %s, Employee ID: %s I earn %.2f. \n", emp.name, emp.empId, emp.salary )
}

func struct_embedding() {
	emp := Employee{
		Person1: Person1{name: "Gurveer", age: 21},
		empId:  "E001", salary: 35000,
	}

	fmt.Println("Name: ", emp.name)
	fmt.Println("Age: ", emp.age)
	fmt.Println("Emp ID: ", emp.empId)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()
}
