package mathematics

import (
	"testing"
)

type testOperation struct {
	a int
	b int
	result int
}

func TestAdd(t *testing.T) {
	var tests = []testOperation{
		{2, 3, 5},
		{-3, 2, -1},
	}

	for _, test := range tests {
		if Add(test.a, test.b) != test.result {
			t.Errorf("Add(%d, %d) = %d, want %d", test.a, test.b, Add(test.a, test.b), test.result)
		}
	}
}

func TestSubstract(t *testing.T) {
	var tests = []testOperation{
		{2, 3, -1},
		{-3, 2, -5},
	}

	for _, test := range tests {
		if Substract(test.a, test.b) != test.result {
			t.Errorf("Substract(%d, %d) = %d, want %d", test.a, test.b, Substract(test.a, test.b), test.result)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []testOperation{
		{2, 3, 6},
		{-3, 2, -6},
	}

	for _, test := range tests {
		if Multiply(test.a, test.b) != test.result {
			t.Errorf("Multiply(%d, %d) = %d, want %d", test.a, test.b, Multiply(test.a, test.b), test.result)
		}
	}
}

func TestDivide(t *testing.T) {
	var tests = []testOperation{
		{4, 2, 2},
		{-3, -1, 3},
	}

	for _, test := range tests {
		result, _ := Divide(test.a, test.b)
		if  result != test.result {
			t.Errorf("Divide(%d, %d) = %d, want %d", test.a, test.b, result, test.result)
		}
	}
}