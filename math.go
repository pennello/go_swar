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

// Average returns the (floor of the) average of x and y.
//
// Commentary from The Aggregate:
//
// This is actually an extension of the "well known" fact that for
// binary integer values x and y, (x+y) equals ((x&y)+(x|y)) equals
// ((x^y)+2*(x&y)).
//
// Given two integer values x and y, the (floor of the) average normally
// would be computed by (x+y)/2; unfortunately, this can yield incorrect
// results due to overflow.  A very sneaky alternative is to use
// (x&y)+((x^y)/2).  The benefit is that this code sequence cannot
// overflow.
//
// Shifts in Go are signed, so this can be simplified to
// (x&y)+((x^y)>>1).
func Average(x, y uint) uint {
	return (x & y) + ((x ^ y) >> 1)
}

// DivCeil32 returns the ceiling of the quotient of a and b.
//
// Uses 64-bit integers.  We take the simple approach here and avoid
// integer overflow by using twice as many bits are are necessary.
//
// Commentary from The Aggregate:
//
// This trick also works if divide is implemented in less obvious ways,
// such as shifts or shift-and-subtract sequences.
func DivCeil32(a, b uint32) uint32 {
	a64 := uint64(a)
	b64 := uint64(b)
	return uint32((a64 + b64 - 1) / b64)
}

// DivRoundNearest32 returns the quotient of a and b, rounded to the
// nearest integer.
//
// Uses 64-bit integers.  We take the simple approach here and avoid
// integer overflow by using twice as many bits are are necessary.
//
// This trick also works if divide is implemented in less obvious ways,
// such as shifts or shift-and-subtract sequences.
func DivRoundNearest32(a, b uint32) uint32 {
	a64 := uint64(a)
	b64 := uint64(b)
	return uint32((a64 + (b64 / 2)) / b64)
}
