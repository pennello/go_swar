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

// Ms1b extracts the most significant 1 bit from 32-bit x.
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the most significant 1 bit (highest
// numbered element of a bit set) can be computed using a SWAR algorithm
// that recursively "folds" the upper bits into the lower bits.  This
// process yields a bit vector with the same most significant 1 as x,
// but all 1's below it.  Bitwise AND of the original value with the
// complement of the "folded" value shifted down by one yields the most
// significant bit.
func Ms1b32(x uint32) uint32 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x & ^(x >> 1)
}

// Ms1b extracts the most significant 1 bit from 64-bit x.
//
// Commentary from The Aggregate:
//
// Given a binary integer value x, the most significant 1 bit (highest
// numbered element of a bit set) can be computed using a SWAR algorithm
// that recursively "folds" the upper bits into the lower bits.  This
// process yields a bit vector with the same most significant 1 as x,
// but all 1's below it.  Bitwise AND of the original value with the
// complement of the "folded" value shifted down by one yields the most
// significant bit.
func Ms1b64(x uint64) uint64 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return x & ^(x >> 1)
}

// Ones8 returns the number of 1 bits in 8-bit x.  This is also known
// as the population count or Hamming weight.
//
// There are many possible implementations of this.  See more here:
//
// http://en.wikipedia.org/wiki/Hamming_weight#Efficient_implementation
func Ones8(x uint8) uint8 {
	x = x&0x55 + ((x >> 1) & 0x55)
	x = x&0x33 + ((x >> 2) & 0x33)
	x = x&0x0f + ((x >> 4) & 0x0f)
	return x
}

// Ones16 returns the number of 1 bits in 16-bit x.  This is also known
// as the population count or Hamming weight.
//
// There are many possible implementations of this.  See more here:
//
// http://en.wikipedia.org/wiki/Hamming_weight#Efficient_implementation
func Ones16(x uint16) uint16 {
	x = x&0x5555 + ((x >> 1) & 0x5555)
	x = x&0x3333 + ((x >> 2) & 0x3333)
	x = x&0x0f0f + ((x >> 4) & 0x0f0f)
	x = x&0x00ff + ((x >> 8) & 0x00ff)
	return x
}

// Ones32 returns the number of 1 bits in 32-bit x.  This is also known
// as the population count or Hamming weight.
//
// There are many possible implementations of this.  See more here:
//
// http://en.wikipedia.org/wiki/Hamming_weight#Efficient_implementation
func Ones32(x uint32) uint32 {
	x = x&0x55555555 + ((x >> 1) & 0x55555555)
	x = x&0x33333333 + ((x >> 2) & 0x33333333)
	x = x&0x0f0f0f0f + ((x >> 4) & 0x0f0f0f0f)
	x = x&0x00ff00ff + ((x >> 8) & 0x00ff00ff)
	x = x&0x0000ffff + ((x >> 16) & 0x0000ffff)
	return x
}

// Ones64 returns the number of 1 bits in 64-bit x.  This is also known
// as the population count or Hamming weight.
//
// There are many possible implementations of this.  See more here:
//
// http://en.wikipedia.org/wiki/Hamming_weight#Efficient_implementation
func Ones64(x uint64) uint64 {
	x = x&0x5555555555555555 + ((x >> 1) & 0x5555555555555555)
	x = x&0x3333333333333333 + ((x >> 2) & 0x3333333333333333)
	x = x&0x0f0f0f0f0f0f0f0f + ((x >> 4) & 0x0f0f0f0f0f0f0f0f)
	x = x&0x00ff00ff00ff00ff + ((x >> 8) & 0x00ff00ff00ff00ff)
	x = x&0x0000ffff0000ffff + ((x >> 16) & 0x0000ffff0000ffff)
	x = x&0x00000000ffffffff + ((x >> 32) & 0x00000000ffffffff)
	return x
}

// Lzc32 returns the number of leading zeroes in 32-bit x.
func Lzc32(x uint32) uint32 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return 32 - Ones32(x)
}

// Lzc64 returns the number of leading zeroes in 64-bit x.
func Lzc64(x uint64) uint64 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return 64 - Ones64(x)
}

// Tzc32 returns the trailing zero count of 32-bit x.
//
// This combines the technique from the least significant 1 bit with the
// population count algorithm.
func Tzc32(x uint32) uint32 {
	// Compare with Ls1b.
	return Ones32((x & -x) - 1)
}

// Tzc64 returns the trailing zero count of 64-bit x.
//
// This combines the technique from the least significant 1 bit with the
// population count algorithm.
func Tzc64(x uint64) uint64 {
	// Compare with Ls1b.
	return Ones64((x & -x) - 1)
}
