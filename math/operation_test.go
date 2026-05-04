package mathematics

import (
	"testing"
)

type testOperation struct {
	a float64
	b float64
	result float64
}

var eps float64 = 0.01

func TestAdd(t *testing.T) {
	var tests = []testOperation{
		{2, 3, 5},
		{-3, 2, -1},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if !IsAlmostEqual(result, test.result, eps) {
			t.Errorf("Add(%f, %f) = %f, want %f", test.a, test.b, result, test.result)
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []testOperation{
		{2, 3, -1},
		{-3, 2, -5},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if !IsAlmostEqual(result, test.result, eps) {
			t.Errorf("Subtract(%f, %f) = %f, want %f", test.a, test.b, result, test.result)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []testOperation{
		{2, 3, 6},
		{-3, 2, -6},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if !IsAlmostEqual(result, test.result, eps) {
			t.Errorf("Multiply(%f, %f) = %f, want %f", test.a, test.b, result, test.result)
		}
	}
}

func TestDivide(t *testing.T) {
	var tests = []testOperation{
		{4, 2, 2},
		{-3, -1, 3},
	}

	for _, test := range tests {
		result, err := Divide(test.a, test.b)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !IsAlmostEqual(result, test.result, eps) {
			t.Errorf("Divide(%f, %f) = %f, want %f", test.a, test.b, result, test.result)
		}
	}
}