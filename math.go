// chris 052315

package swar

import (
	"math"
)

// Abs32 returns the absolute value of a 32-bit float.
//
// Special cases are:
//	Abs32(±Inf) = Inf
//	Abs32(NaN) = NaN
func Abs32(x float32) float32 {
	xb := math.Float32bits(x)
	return math.Float32frombits(xb & 0x7fffffff)
}

// Abs64 returns the absolute value of a 64-bit float.
//
// Special cases are:
//	Abs64(±Inf) = Inf
//	Abs64(NaN) = NaN
func Abs64(x float64) float64 {
	xb := math.Float64bits(x)
	return math.Float64frombits(xb & 0x7fffffffffffffff)
}
