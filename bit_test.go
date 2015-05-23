// chris 052315

package swar

import (
	"strconv"
	"strings"
	"testing"

	"math/rand"

	fixrand "chrispennello.com/go/util/fix/math/rand"
)

// Maybe move this into util/fix/strings?
func reverseBytes(s string) string {
	L := len(s)
	sb := make([]byte, L)
	for i := 0; i < L; i++ {
		sb[i] = byte(s[L-i-1])
	}
	return string(sb)
}

func testReverseBytes(t *testing.T, s, check string) {
	out := reverseBytes(s)
	if out != check {
		t.Errorf("ReverseBytes(%v) != %v (is %v)", s, check, out)
	}
}

func TestReverseBytes(t *testing.T) {
	testReverseBytes(t, "", "")
	testReverseBytes(t, "a", "a")
	testReverseBytes(t, "ab", "ba")
	testReverseBytes(t, "abc", "cba")
}

func reverseBits32Ref(x uint32) uint32 {
	xs := strconv.FormatUint(uint64(x), 2)
	xr := reverseBytes(xs)
	xr += strings.Repeat("0", 32-len(xr))
	xx, err := strconv.ParseUint(xr, 2, 32)
	if err != nil {
		panic(err)
	}
	return uint32(xx)
}

func testReverseBits32(t *testing.T, x uint32) {
	out := ReverseBits32(x)
	check := reverseBits32Ref(x)
	if out != check {
		t.Errorf("ReverseBits32(%b) != %b (is %b)", x, check, out)
	}
}

func TestReverseBits32(t *testing.T) {
	testReverseBits32(t, 0)
	testReverseBits32(t, 1)
	testReverseBits32(t, 0xffffffff)
	testReverseBits32(t, 0xffff0000)

	for i := 0; i < 1000; i++ {
		testReverseBits32(t, rand.Uint32())
	}
}

func reverseBits64Ref(x uint64) uint64 {
	xs := strconv.FormatUint(x, 2)
	xr := reverseBytes(xs)
	xr += strings.Repeat("0", 64-len(xr))
	xx, err := strconv.ParseUint(xr, 2, 64)
	if err != nil {
		panic(err)
	}
	return xx
}

func testReverseBits64(t *testing.T, x uint64) {
	out := ReverseBits64(x)
	check := reverseBits64Ref(x)
	if out != check {
		t.Errorf("ReverseBits64(%b) != %b (is %b)", x, check, out)
	}
}

func TestReverseBits64(t *testing.T) {
	testReverseBits64(t, 0)
	testReverseBits64(t, 1)
	testReverseBits64(t, 0xffffffff)
	testReverseBits64(t, 0xffff0000)

	testReverseBits64(t, 0xffffffffffffffff)
	testReverseBits64(t, 0xffffffff00000000)

	for i := 0; i < 1000; i++ {
		testReverseBits64(t, fixrand.Uint64())
	}
}
