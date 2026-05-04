package mathematics

import (
	"fmt"
	"strconv"
	"math"
)

// Returns two float64 numbers and an error given an array of 2 strings
//
// Examples:
//
// ParseTwoOperands({"12", "3"}) = 12, 3, nil
func ParseTwoOperands(args []string) (float64, float64, error) {
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("Unexpected Number of arguments: want 2, got %d", len(args))
	}
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid first argument %q: %v", args[0], err)
	}
	b, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid second argument %q: %v", args[1], err)
	}
	return a, b, nil
} 

// Returns whether two float64 numbers are almost equal given a margin eps
// 
// Examples:
//
// IsAlmostEqual(10.022, 10.031, 0.01) = true
func IsAlmostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}