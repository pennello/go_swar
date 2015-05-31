// chris 052315

package swar

// IntegerSelect32 implements a branchless, lookup-free version of
// a < b ? c : d.
//
// Uses 64-bit integers.  We take the simple approach here and avoid
// integer over- and underflow by using twice as many bits are are
// necessary.
//
// Commentary from The Aggregate:
//
// Logically, this works because the shift by (WORDBITS-1) replicates
// the sign bit to create a mask.
func IntegerSelect32(a, b, c, d int32) int32 {
	a64 := int64(a)
	b64 := int64(b)
	c64 := int64(c)
	d64 := int64(d)
	return int32((((a64 - b64) >> (64 - 1)) & (c64 ^ d64)) ^ d64)
}
