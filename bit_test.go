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

func ls1bRef(x uint) uint {
	if x == 0 {
		return 0
	}
	for place := uint(0); ; place++ {
		b := uint(1) << place
		if b&x == b {
			return b
		}
	}
}

func testLs1b(t *testing.T, x uint) {
	out := Ls1b(x)
	check := ls1bRef(x)
	if out != check {
		t.Errorf("Ls1b(%v) != %v (is %v)", x, check, out)
	}
}

func TestLs1b(t *testing.T) {
	testLs1b(t, 0)
	testLs1b(t, 1)
	testLs1b(t, 2)
	testLs1b(t, 0x80000000)
	testLs1b(t, 0x80000001)
	testLs1b(t, 0x80000010)
	testLs1b(t, 0xffffffff)

	testLs1b(t, 0x8000000000000000)
	testLs1b(t, 0x8000000000000001)
	testLs1b(t, 0x8000000000000010)
	testLs1b(t, 0xffffffffffffffff)

	for i := 0; i < 1000; i++ {
		testLs1b(t, uint(fixrand.Uint64()))
	}
}

func ms1b32Ref(x uint32) uint32 {
	if x == 0 {
		return 0
	}
	for place := uint(32 - 1); ; place-- {
		b := uint32(1) << place
		if b&x == b {
			return b
		}
	}
}

func testMs1b32(t *testing.T, x uint32) {
	out := Ms1b32(x)
	check := ms1b32Ref(x)
	if out != check {
		t.Errorf("Ms1b32(%v) != %v (is %v)", x, check, out)
	}
}

func TestMs1b32(t *testing.T) {
	testMs1b32(t, 0)
	testMs1b32(t, 1)
	testMs1b32(t, 2)
	testMs1b32(t, 0xffffffff)
	testMs1b32(t, 0x80000000)

	for i := 0; i < 1000; i++ {
		testMs1b32(t, rand.Uint32())
	}
}

func ms1b64Ref(x uint64) uint64 {
	if x == 0 {
		return 0
	}
	for place := uint(64 - 1); ; place-- {
		b := uint64(1) << place
		if b&x == b {
			return b
		}
	}
}

func testMs1b64(t *testing.T, x uint64) {
	out := Ms1b64(x)
	check := ms1b64Ref(x)
	if out != check {
		t.Errorf("Ms1b64(%v) != %v (is %v)", x, check, out)
	}
}

func TestMs1b64(t *testing.T) {
	testMs1b64(t, 0)
	testMs1b64(t, 1)
	testMs1b64(t, 2)
	testMs1b64(t, 0xffffffff)
	testMs1b64(t, 0x80000000)

	for i := 0; i < 1000; i++ {
		testMs1b64(t, fixrand.Uint64())
	}
}

func ones8Ref(x uint8) uint8 {
	ones := uint8(0)
	for place := uint(0); place < 8; place++ {
		b := uint8(1) << place
		if x&b == b {
			ones++
		}
	}
	return ones
}

func testOnes8(t *testing.T, x uint8) {
	out := Ones8(x)
	check := ones8Ref(x)
	if out != check {
		xs := strconv.FormatUint(uint64(x), 2)
		t.Errorf("Ones8(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestOnes8(t *testing.T) {
	testOnes8(t, 0)
	testOnes8(t, 1)
	testOnes8(t, 2)
	testOnes8(t, 3)
	testOnes8(t, 4)
	testOnes8(t, 5)
	testOnes8(t, 0xff)

	for i := 0; i < 1000; i++ {
		testOnes8(t, uint8(rand.Uint32()))
	}
}

func ones16Ref(x uint16) uint16 {
	ones := uint16(0)
	for place := uint(0); place < 16; place++ {
		b := uint16(1) << place
		if x&b == b {
			ones++
		}
	}
	return ones
}

func testOnes16(t *testing.T, x uint16) {
	out := Ones16(x)
	check := ones16Ref(x)
	if out != check {
		xs := strconv.FormatUint(uint64(x), 2)
		t.Errorf("Ones16(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestOnes16(t *testing.T) {
	testOnes16(t, 0)
	testOnes16(t, 1)
	testOnes16(t, 2)
	testOnes16(t, 3)
	testOnes16(t, 4)
	testOnes16(t, 5)
	testOnes16(t, 0xffff)

	for i := 0; i < 1000; i++ {
		testOnes16(t, uint16(rand.Uint32()))
	}
}

func ones32Ref(x uint32) uint32 {
	ones := uint32(0)
	for place := uint(0); place < 32; place++ {
		b := uint32(1) << place
		if x&b == b {
			ones++
		}
	}
	return ones
}

func testOnes32(t *testing.T, x uint32) {
	out := Ones32(x)
	check := ones32Ref(x)
	if out != check {
		xs := strconv.FormatUint(uint64(x), 2)
		t.Errorf("Ones32(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestOnes32(t *testing.T) {
	testOnes32(t, 0)
	testOnes32(t, 1)
	testOnes32(t, 2)
	testOnes32(t, 3)
	testOnes32(t, 4)
	testOnes32(t, 5)
	testOnes32(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testOnes32(t, rand.Uint32())
	}
}

func ones64Ref(x uint64) uint64 {
	ones := uint64(0)
	for place := uint(0); place < 64; place++ {
		b := uint64(1) << place
		if x&b == b {
			ones++
		}
	}
	return ones
}

func testOnes64(t *testing.T, x uint64) {
	out := Ones64(x)
	check := ones64Ref(x)
	if out != check {
		xs := strconv.FormatUint(x, 2)
		t.Errorf("Ones64(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestOnes64(t *testing.T) {
	testOnes64(t, 0)
	testOnes64(t, 1)
	testOnes64(t, 2)
	testOnes64(t, 3)
	testOnes64(t, 4)
	testOnes64(t, 5)
	testOnes64(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testOnes64(t, fixrand.Uint64())
	}
}

func lzc32Ref(x uint32) uint32 {
	place := uint(32 - 1)
	lz := uint32(0)
	for {
		b := uint32(1) << place
		if b & ^x == b {
			lz++
		} else {
			break
		}
		if place == 0 {
			break
		}
		place--
	}
	return lz
}

func testLzc32(t *testing.T, x uint32) {
	out := Lzc32(x)
	check := lzc32Ref(x)
	if out != check {
		xs := strconv.FormatUint(uint64(x), 2)
		t.Errorf("Lzc32(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestLzc32(t *testing.T) {
	testLzc32(t, 0)
	testLzc32(t, 1)
	testLzc32(t, 2)
	testLzc32(t, 3)
	testLzc32(t, 4)
	testLzc32(t, 5)
	testLzc32(t, 0xffffffff)
	testLzc32(t, 0x80000000)
	testLzc32(t, 0x8000)

	for i := 0; i < 1000; i++ {
		testLzc32(t, rand.Uint32())
	}
}

func lzc64Ref(x uint64) uint64 {
	place := uint(64 - 1)
	lz := uint64(0)
	for {
		b := uint64(1) << place
		if b & ^x == b {
			lz++
		} else {
			break
		}
		if place == 0 {
			break
		}
		place--
	}
	return lz
}

func testLzc64(t *testing.T, x uint64) {
	out := Lzc64(x)
	check := lzc64Ref(x)
	if out != check {
		xs := strconv.FormatUint(x, 2)
		t.Errorf("Lzc64(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestLzc64(t *testing.T) {
	testLzc64(t, 0)
	testLzc64(t, 1)
	testLzc64(t, 2)
	testLzc64(t, 3)
	testLzc64(t, 4)
	testLzc64(t, 5)
	testLzc64(t, 0xffffffff)
	testLzc64(t, 0x80000000)
	testLzc64(t, 0x8000)

	testLzc64(t, 0xffffffffffffffff)
	testLzc64(t, 0x8000000000000000)
	testLzc64(t, 0x80000000)

	for i := 0; i < 1000; i++ {
		testLzc64(t, fixrand.Uint64())
	}
}

func tzc32Ref(x uint32) uint32 {
	tzc := uint32(0)
	for place := uint(0); place < 32; place++ {
		b := uint32(1) << place
		if b & ^x == b {
			tzc++
		} else {
			break
		}
	}
	return tzc
}

func testTzc32(t *testing.T, x uint32) {
	out := Tzc32(x)
	check := tzc32Ref(x)
	if out != check {
		xs := strconv.FormatUint(uint64(x), 2)
		t.Errorf("Tzc32(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestTzc32(t *testing.T) {
	testTzc32(t, 0)
	testTzc32(t, 1)
	testTzc32(t, 2)
	testTzc32(t, 3)
	testTzc32(t, 4)
	testTzc32(t, 0x80000000)
	testTzc32(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testTzc32(t, rand.Uint32())
	}
}

func tzc64Ref(x uint64) uint64 {
	tzc := uint64(0)
	for place := uint(0); place < 64; place++ {
		b := uint64(1) << place
		if b & ^x == b {
			tzc++
		} else {
			break
		}
	}
	return tzc
}

func testTzc64(t *testing.T, x uint64) {
	out := Tzc64(x)
	check := tzc64Ref(x)
	if out != check {
		xs := strconv.FormatUint(x, 2)
		t.Errorf("Tzc64(0b%v) != %v (is %v)", xs, check, out)
	}
}

func TestTzc64(t *testing.T) {
	testTzc64(t, 0)
	testTzc64(t, 1)
	testTzc64(t, 2)
	testTzc64(t, 3)
	testTzc64(t, 4)
	testTzc64(t, 0x80000000)
	testTzc64(t, 0xffffffff)

	testTzc64(t, 0x8000000000000000)
	testTzc64(t, 0xffffffffffffffff)

	for i := 0; i < 1000; i++ {
		testTzc64(t, fixrand.Uint64())
	}
}
