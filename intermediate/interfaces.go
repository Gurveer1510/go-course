package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rect) perimeter() float64 {
	return 2 * r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func interfaces() {
	r := rect{width: 4.23,height: 9.32}
	c := circle{radius: 3.14}

	measure(r)
	measure(c)

	myPrinter(1, "John", 45.66, true)
	printYpe(12)
	printYpe("Gurveer")
}

func printYpe(i interface{}) {
	switch i.(type) {
	case int :
		fmt.Println("Type: int")
	case string :
		fmt.Println("Type: string")
	case float64:
		fmt.Println("Type: float64")
	case bool :
		fmt.Println("Type: bool")
	default:
		fmt.Println("Unknown Type")
	}
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}