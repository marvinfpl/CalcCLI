package mathematics

import (
	"math"
)

func Cos(a float64) float64 {
	return math.Cos(a)
}

func Sin(a float64) float64 {
	return math.Sin(a)
}

func Exp(a float64) float64 {
	return math.Exp(a)
}

func Log(a float64, basis int) float64 {
	if basis == 0 {
		return math.Log(a)
	}
	x := math.Log(a) / math.Log((float64)(basis))
	return x
}

