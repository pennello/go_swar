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
// integer overflow by using twice as many bits as are necessary.
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
// integer overflow by using twice as many bits as are necessary.
//
// This trick also works if divide is implemented in less obvious ways,
// such as shifts or shift-and-subtract sequences.
func DivRoundNearest32(a, b uint32) uint32 {
	a64 := uint64(a)
	b64 := uint64(b)
	return uint32((a64 + (b64 / 2)) / b64)
}

// Min32 returns the minimum of x and y.
//
// Uses 64-bit integers.  We take the simple approach here and avoid
// integer over- and underflow by using twice as many bits are are
// necessary.
//
// Commentary from The Aggregate:
//
// Logically, this works because the shift by (WORDBITS-1) replicates
// the sign bit to create a mask
func Min32(x, y int32) int32 {
	x64 := int64(x)
	y64 := int64(y)
	return int32(x64 + (((y64 - x64) >> (64 - 1)) & (y64 - x64)))
}

// Max32 returns the maximum of x and y.
//
// Uses 64-bit integers.  We take the simple approach here and avoid
// integer over- and underflow by using twice as many bits are are
// necessary.
//
// Commentary from The Aggregate:
//
// Logically, this works because the shift by (WORDBITS-1) replicates
// the sign bit to create a mask
func Max32(x, y int32) int32 {
	x64 := int64(x)
	y64 := int64(y)
	return int32(x64 - (((x64 - y64) >> (64 - 1)) & (x64 - y64)))
}

// IsPow2 returns whether x is a power of 2.
//
// Special case is:
//
//	IsPow2(0) = 0
func IsPow2(x uint) bool {
	return x&(x-1) == 0
}

// Nlpo232 returns the next largest power of 2 from 32-bit x.
//
// Special cases are:
//
//	Nlpo232(x) = 0 // where x has the high bit set
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the next largest power of 2 can be
// computed by a SWAR algorithm that recursively "folds" the upper bits
// into the lower bits.  This process yields a bit vector with the same
// most significant 1 as x, but all 1's below it.  Adding 1 to that
// value yields the next largest power of 2.
func Nlpo232(x uint32) uint32 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x + 1
}

// Nlpo264 returns the next largest power of 2 from 64-bit x.
//
// Special cases are:
//
//	Nlpo264(x) = 0 // where x has the high bit set
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the next largest power of 2 can be
// computed by a SWAR algorithm that recursively "folds" the upper bits
// into the lower bits.  This process yields a bit vector with the same
// most significant 1 as x, but all 1's below it.  Adding 1 to that
// value yields the next largest power of 2.
func Nlpo264(x uint64) uint64 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return x + 1
}

// SameWithinTolerance32 determines if 32-bit integers a and b have the
// same value within the given 32-bit tolerance c.
//
// Uses 64-bit integers.  We take the simple approach here and avoid
// integer over- and underflow by using twice as many bits are are
// necessary.
//
// Commentary from The Aggregate:
//
// The obvious test would be something like ((a>b)?(a-b):(b-a))<c, which
// isn't horrifically inefficient, but does involve a conditional
// branch. Alternatively, abs(a-b)<c would do... but again, it takes
// some cleverness to implement abs() without a conditional jump. Here's
// a branchless alternative.
//
// If (a-b)>0, then (b-a)<0; similarly, if (a-b)<0, then (b-a)>0. Both
// can't be greater than 0 simultaneously. Suppose that (a-b)>0.
// Subtracting ((a-b)-c) will produce a negative result iff a and b are
// within c of each other. Of course, our assumption requires (b-a)<0,
// so ((b-a)-c) simply becomes more negative (assuming the value doesn't
// wrap around). Generalizing, if either ((a-b)-c)>0 or ((b-a)-c)>0 then
// the values of a and b are not the same within tolerance c. In other
// words, they are within tolerance if:
//
//	(((a-b-c)&(b-a-c))<0)
//
// This test can be rewritten a variety of ways. The <0 part is really
// just examining the sign bit, so a mask or shift could be used to
// extract the bit value instead. For example, using 32-bit words,
// (((a-b-c)&(b-a-c))>>31) using unsigned >> will produce the value 1
// for true or 0 for false. It is also possible to factor-out t=a-b,
// giving:
//
//	(((t-c)&(-t-c))<0)
//
// Which is really equivalent to abs(t)<c.
func SameWithinTolerance32(a, b, c int32) bool {
	a64 := int64(a)
	b64 := int64(b)
	c64 := int64(c)
	t64 := a64 - b64
	return ((t64-c64)&(-t64-c64)>>(64-1))&1 == 1
}

// Log2Floor32 returns the floor of the base 2 logarithm of 32-bit x.
//
// Special case is:
//
//	Log2Floor32(0) = 0
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the floor of the base 2 log of that
// number efficiently can be computed by the application of two
// variable-precision SWAR algorithms. The first "folds" the upper bits
// into the lower bits to construct a bit vector with the same most
// significant 1 as x, but all 1's below it. The second SWAR algorithm
// is population count, defined elsewhere in this package.
func Log2Floor32(x uint32) uint32 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return Ones32(x >> 1)
}

// Log2Floor64 returns the floor of the base 2 logarithm of 64-bit x.
//
// Special case is:
//
//	Log2Floor64(0) = 0
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the floor of the base 2 log of that
// number efficiently can be computed by the application of two
// variable-precision SWAR algorithms. The first "folds" the upper bits
// into the lower bits to construct a bit vector with the same most
// significant 1 as x, but all 1's below it. The second SWAR algorithm
// is population count, defined elsewhere in this package.
func Log2Floor64(x uint64) uint64 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return Ones64(x >> 1)
}

// Log2Ceil32 returns the ceiling of the base 2 logarithm of 32-bit x.
//
// Special case is:
//
//	Log2Ceil32(0) = 0
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the floor of the base 2 log of that
// number efficiently can be computed by the application of two
// variable-precision SWAR algorithms. The first "folds" the upper bits
// into the lower bits to construct a bit vector with the same most
// significant 1 as x, but all 1's below it. The second SWAR algorithm
// is population count, defined elsewhere in this package.
//
// Suppose instead that you want the ceiling of the base 2 log. The
// floor and ceiling are identical if x is a power of two; otherwise,
// the result is 1 too small. This can be corrected using the power of 2
// test followed with the comparison-to-mask shift used in integer
// minimum/maximum.
func Log2Ceil32(x uint32) uint32 {
	y := x & (x - 1) // Like IsPow2.
	y |= -y
	y >>= 32 - 1
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return Ones32(x>>1) + y
}

// Log2Ceil64 returns the ceiling of the base 2 logarithm of 64-bit x.
//
// Special case is:
//
//	Log2Ceil64(0) = 0
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the floor of the base 2 log of that
// number efficiently can be computed by the application of two
// variable-precision SWAR algorithms. The first "folds" the upper bits
// into the lower bits to construct a bit vector with the same most
// significant 1 as x, but all 1's below it. The second SWAR algorithm
// is population count, defined elsewhere in this package.
//
// Suppose instead that you want the ceiling of the base 2 log. The
// floor and ceiling are identical if x is a power of two; otherwise,
// the result is 1 too small. This can be corrected using the power of 2
// test followed with the comparison-to-mask shift used in integer
// minimum/maximum.
func Log2Ceil64(x uint64) uint64 {
	y := x & (x - 1) // Like IsPow2.
	y |= -y
	y >>= 64 - 1
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return Ones64(x>>1) + y
}
