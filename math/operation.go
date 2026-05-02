package mathematics

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		fmt.Println("Cannot divide by 0")
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}