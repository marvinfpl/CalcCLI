package mathematics

import (
	"testing"
	"math"
)

type testFunction struct {
	a float64
	result float64
}

func TestCos(t *testing.T) {
	var tests = []testFunction{
		{0.0, 1.0},
		{90.0, 0.0},
	}

	for _, test := range tests {
		if Cos(test.a) != test.result {
			t.Errorf("Cos(%f) = %f want %f", test.a, Cos(test.a), test.result)
		}
	}
}

func TestSin(t *testing.T) {
	var tests = []testFunction{
		{0.0, 0.0},
		{90.0, 1.0},
	}

	for _, test := range tests {
		if Sin(test.a) != test.result {
			t.Errorf("Sin(%f) = %f want %f", test.a, Sin(test.a), test.result)
		}
	}
}

func TestExp(t *testing.T) {
	var tests = []testFunction{
		{0.0, 1.0},
		{1.0, math.Exp(1)},
	}

	for _, test := range tests {
		if Exp(test.a) != test.result {
			t.Errorf("Exp(%f) = %f want %f", test.a, Exp(test.a), test.result)
		}
	}
}

func TestLog(t *testing.T) {
	var tests = []struct{
		a float64
		basis int
		result float64
	}{
		{1.0, 0, 0.0,},
		{math.Exp(1), 0, 1.0},
	}

	for _, test := range tests {
		if Log(test.a, 0) != test.result {
			t.Errorf("Log(%f) = %f want %f", test.a, Log(test.a, test.basis), test.result)
		}
	}
}