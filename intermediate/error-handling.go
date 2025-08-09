package intermediate

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: square root of negative number")
	}
	res := math.Sqrt(x)
	return res, nil
}

type myError struct {
	message string
}

func (err *myError) Error() string {
	return fmt.Sprintf("%s", err.message)
}

func eProcess() error {
	return &myError{"Custom error message"}
}

type customError struct {
	code    int
	message string
	err     error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error: %d: %s, %v\n", e.code, e.message, e.err)
}

func doSomething() error {
	err := donSomethingElse()
	if err != nil {
		return &customError{
			code: 500,
			message: "Something went wrong",
			err: err,
		}
	}
	return nil
}

func donSomethingElse() error {
	return errors.New("internal error")
}


func error_handling() {
	// res, err := sqrt(-1)

	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	return
	// }
	// fmt.Println(res)

	// err1 := eProcess()
	// if err1 != nil {
	// 	fmt.Println("Error: ", err1)
	// 	return
	// }

	res := doSomething()
	if res != nil {
		fmt.Println(res)
	}
}
