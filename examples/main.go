package main

import (
	"fmt"
)

// divideByZeroError is a custom error type
type divideByZeroError struct{}

func (e *divideByZeroError) Error() string {
	return "cannot divide by zero"
}

// divide divides two integers and returns the result or an error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &divideByZeroError{}
	}
	return a / b, nil
}

func main() {
	a, b := 10, 0

	result, err := divide(a, b)
	if err != nil {
		// Checking the specific error type using type assertion
		if e, ok := err.(*divideByZeroError); ok {
			fmt.Println("Error:", e.Error())
		} else {
			fmt.Println("Unknown Error:", err.Error())
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}
