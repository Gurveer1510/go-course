package intermediate

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func methods() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	fmt.Println("Area of 10x9 Rectanlge:", rect.Area())
	rect.Scale(10)
	fmt.Println("Area of 10x9 Rectanlge:", rect.Area())

	num1 := MyInt(-2)
	num2 := MyInt(2)
	fmt.Println(num1.isPositive())
	fmt.Println(num2.isPositive())
	fmt.Println(num1.WelcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 3, width: 5}}

	fmt.Println(s.Area())

}

type MyInt int

func (m MyInt) isPositive() bool {
	return m > 0
}

func (MyInt) WelcomeMessage() string {
	return "Welcome to MyInt Type"
}
