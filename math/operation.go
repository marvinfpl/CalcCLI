package mathematics

import (
	"errors"
)

var ErrDivisionByZero error = errors.New("Cannot divide by 0")

// Returns the addition of 2 real numbers.
//
// Special cases: 
//
// Add(-Inf, +Inf) = NaN
//
// Add(NaN, x) = NaN
func Add(a, b float64) float64 {
	return a + b
}

// Returns the subtraction of 2 real numbers.
//
// Special cases:
//
// Subtract(-Inf, +Inf) = NaN
//
// Subtract(NaN, x) = NaN
func Subtract(a, b float64) float64 {
	return a - b
}

// Returns the multiplication of 2 real numbers.
//
// Special cases:
//
// Multiply(0, ±Inf) = NaN
//
// Multiply(NaN, x) = NaN
func Multiply(a, b float64) float64 {
	return a * b
}

// Returns the division of 2 real numbers or an error if Divide(a, 0) is called.
//
// Special cases:
//
// Divide(±Inf, ±Inf) = NaN
//
// Divide(0, 0) = NaN
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}