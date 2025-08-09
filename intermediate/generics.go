package intermediate

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func ( s *Stack[T]) Push(ele T) {
	s.elements = append(s.elements, ele)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	ele := s.elements[len(s.elements) - 1]
	s.elements = s.elements[:len(s.elements)-1]
	return ele, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) PrintAll() {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Println("Stack Contents: ")
	for _, v := range s.elements {
		fmt.Println(v)
	}
}

func generics() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	x1, y1 := "John", "Doe"
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, y1)

	s := Stack[int]{elements: []int{}}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	s.PrintAll()

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	s.PrintAll()
	fmt.Println(s.isEmpty())
}
