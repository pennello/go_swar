// chris 052315

package swar

// ReverseBits32 returns a 32-bit integer with the bits from x reversed.
func ReverseBits32(x uint32) uint32 {
	x = ((x & 0xaaaaaaaa) >> 1) | ((x & 0x55555555) << 1)
	x = ((x & 0xcccccccc) >> 2) | ((x & 0x33333333) << 2)
	x = ((x & 0xf0f0f0f0) >> 4) | ((x & 0x0f0f0f0f) << 4)
	x = ((x & 0xff00ff00) >> 8) | ((x & 0x00ff00ff) << 8)
	return (x >> 16) | (x << 16)
}

// ReverseBits64 returns a 64-bit integer with the bits from x reversed.
func ReverseBits64(x uint64) uint64 {
	x = ((x & 0xaaaaaaaaaaaaaaaa) >> 1) | ((x & 0x5555555555555555) << 1)
	x = ((x & 0xcccccccccccccccc) >> 2) | ((x & 0x3333333333333333) << 2)
	x = ((x & 0xf0f0f0f0f0f0f0f0) >> 4) | ((x & 0x0f0f0f0f0f0f0f0f) << 4)
	x = ((x & 0xff00ff00ff00ff00) >> 8) | ((x & 0x00ff00ff00ff00ff) << 8)
	x = ((x & 0xffff0000ffff0000) >> 16) | ((x & 0x0000ffff0000ffff) << 16)
	return (x >> 32) | (x << 32)
}

// Ls1b extracts the least significant 1 bit from x.
//
// Commentary from The Aggregate:
//
// This can be useful for extracting the lowest numbered element of a
// bit set.  The reason this works is that it is equivalent to
// (x & ((~x) + 1)); any trailing zero bits in x become ones in ~x,
// adding 1 to that carries into the following bit, and AND with x
// yields only the flipped bit... the original position of the least
// significant 1 bit.
func Ls1b(x uint) uint {
	return x & -x
}
