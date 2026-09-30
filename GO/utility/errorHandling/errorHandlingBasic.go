package errorhandling

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot be divided by zero")
	}
	return a / b, nil
}

func ErrorHandlingBasic() {
	var a int
	fmt.Print("Enter number 1: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Enter number 2: ")
	fmt.Scan(&b)

	result, err := divide(a, b)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Result: %v", result)
}
