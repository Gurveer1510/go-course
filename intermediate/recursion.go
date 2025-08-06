package intermediate

import "fmt"

func recursion() {
	// fmt.Println(Factorial(5))
	// fmt.Println(SumOfDigits(99))
	// fmt.Println(SumOfDigits(9))
	PrintFibonacci(0, 8)
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return n * Factorial(n-1)
}

func PrintFibonacci (i , n int) {
	if i >= n {
		return
	}
	fmt.Println(Fibonacci(i))
	PrintFibonacci(i+1, n)
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func SumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + SumOfDigits(n/10)

}
