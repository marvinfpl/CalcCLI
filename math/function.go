package mathematics

import (
	"math"
)

//Returns the cosine of the argument in radian.
//
//Special cases:
//
//Cos(±Inf) = NaN
//
//Cos(NaN) = NaN
func Cos(a float64) float64 {
	return math.Cos(a)
}

//Returns the sine of the argument in radian.
//
//Special cases:
//
//Sin(±Inf) = NaN
//
//Sin(NaN) = NaN
func Sin(a float64) float64 {
	return math.Sin(a) 
}

//Returns the base e exponent of x (e**x).
//
//Special cases:
//
//Exp(+Inf) = +Inf
//
//Exp(NaN) = NaN
func Exp(a float64) float64 {
	return math.Exp(a)
}

//Returns the logarithm of the argument given the basis.
//
// Log(a, 0) returns the natural logarithm, exponential basis (Log(e, 0) = 1).
//
//Special cases:
//
//Log(+Inf, b) = +Inf
//
//Log(0, b) = -Inf
//
//Log(a < 0, b) = NaN
//
//Log(NaN, b) = NaN
func Log(a float64, basis int) float64 {
	if basis <= 0 {
		return math.Log(a)
	}
	x := math.Log(a) / math.Log((float64)(basis))
	return x
}