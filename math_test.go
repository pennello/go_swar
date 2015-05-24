// chris 052315

package swar

import (
	"math"
	"testing"

	"math/big"
	"math/rand"

	fixmath "chrispennello.com/go/util/fix/math"
	fixrand "chrispennello.com/go/util/fix/math/rand"
)

func testAbs32(t *testing.T, x float32) {
	out := Abs32(x)
	check := float32(math.Abs(float64(x)))
	if out != check {
		t.Errorf("Abs32(%v) != %v (is %v)", x, check, out)
	}
}

func testAbs32Inf(t *testing.T, sign int) {
	x := math.Inf(sign)
	out := Abs32(float32(x))
	if !math.IsInf(float64(out), 1) {
		t.Errorf("Abs32(%v) \"!=\" %v (is %v)", x, x, out)
	}
}

func testAbs32NaN(t *testing.T) {
	x := math.NaN()
	out := Abs32(float32(x))
	if !math.IsNaN(float64(out)) {
		t.Errorf("Abs32(%v) \"!=\" %v (is %v)", x, x, out)
	}
}

func TestAbs32(t *testing.T) {
	testAbs32Inf(t, 1)
	testAbs32Inf(t, -1)
	testAbs32NaN(t)

	testAbs32(t, 0)
	testAbs32(t, .25)
	testAbs32(t, .5)
	testAbs32(t, .75)
	testAbs32(t, 1)

	testAbs32(t, -0)
	testAbs32(t, -.25)
	testAbs32(t, -.5)
	testAbs32(t, -.75)
	testAbs32(t, -1)

	for i := 0; i < 1000; i++ {
		x := math.Float32frombits(rand.Uint32())
		if math.IsNaN(float64(x)) || math.IsInf(float64(x), 0) {
			continue
		}
		testAbs32(t, x)
	}
}

func testAbs64(t *testing.T, x float64) {
	out := Abs64(x)
	check := math.Abs(x)
	if out != check {
		t.Errorf("Abs64(%v) != %v (is %v)", x, check, out)
	}
}

func testAbs64Inf(t *testing.T, sign int) {
	x := math.Inf(sign)
	out := Abs64(x)
	if !math.IsInf(out, 1) {
		t.Errorf("Abs64(%v) \"!=\" %v (is %v)", x, x, out)
	}
}

func testAbs64NaN(t *testing.T) {
	x := math.NaN()
	out := Abs64(x)
	if !math.IsNaN(out) {
		t.Errorf("Abs64(%v) \"!=\" %v (is %v)", x, x, out)
	}
}

func TestAbs64(t *testing.T) {
	testAbs64Inf(t, 1)
	testAbs64Inf(t, -1)
	testAbs64NaN(t)

	testAbs64(t, 0)
	testAbs64(t, .25)
	testAbs64(t, .5)
	testAbs64(t, .75)
	testAbs64(t, 1)

	testAbs64(t, -.25)
	testAbs64(t, -.5)
	testAbs64(t, -.75)
	testAbs64(t, -1)

	for i := 0; i < 1000; i++ {
		x := math.Float64frombits(fixrand.Uint64())
		if math.IsNaN(x) || math.IsInf(x, 0) {
			continue
		}
		testAbs64(t, x)
	}
}

func averageRef(x, y uint) uint {
	// Initially tried using float64s, but they don't have enough precision!
	xb := big.NewInt(0)
	yb := big.NewInt(0)
	sm := big.NewInt(0)
	av := big.NewInt(0)
	xb.SetUint64(uint64(x))
	yb.SetUint64(uint64(y))
	sm.Add(xb, yb)
	av.Div(sm, big.NewInt(2))
	return uint(av.Uint64())
}

func testAverage(t *testing.T, x, y uint) {
	out := Average(x, y)
	check := averageRef(x, y)
	if out != check {
		t.Errorf("Average(0x%x, 0x%x) != 0x%x (is 0x%x)", x, y, check, out)
	}
}

func TestAverage(t *testing.T) {
	testAverage(t, 1, 2)
	testAverage(t, 150, 125)

	for i := 0; i < 1000; i++ {
		x := uint(fixrand.Uint64())
		y := uint(fixrand.Uint64())
		testAverage(t, x, y)
	}
}

func divCeil32Ref(a, b uint32) uint32 {
	af := float64(a)
	bf := float64(b)
	return uint32(math.Ceil(af / bf))
}

func testDivCeil32(t *testing.T, a, b uint32) {
	out := DivCeil32(a, b)
	check := divCeil32Ref(a, b)
	if out != check {
		t.Errorf("DivCeil32(%v, %v) != %v (is %v)", a, b, check, out)
	}
}

func TestDivCeil32(t *testing.T) {
	testDivCeil32(t, 0, 1)
	testDivCeil32(t, 1, 2)
	testDivCeil32(t, 2, 2)
	testDivCeil32(t, 3, 2)
	testDivCeil32(t, 1, 3)
	testDivCeil32(t, 2, 3)
	testDivCeil32(t, 3, 3)
	testDivCeil32(t, 4, 3)

	for i := 0; i < 1000; i++ {
		a := rand.Uint32()
		b := rand.Uint32()
		if b == 0 {
			continue
		}
		testDivCeil32(t, a, b)
	}
}

func divRoundNearest32Ref(a, b uint32) uint32 {
	af := float64(a)
	bf := float64(b)
	return uint32(fixmath.Round(af / bf))
}

func testDivRoundNearest32(t *testing.T, a, b uint32) {
	out := DivRoundNearest32(a, b)
	check := divRoundNearest32Ref(a, b)
	if out != check {
		t.Errorf("DivRoundNearest32(%v, %v) != %v (is %v)", a, b, check, out)
	}
}

func TestDivRoundNearest32(t *testing.T) {
	testDivRoundNearest32(t, 0, 1)
	testDivRoundNearest32(t, 1, 1)
	testDivRoundNearest32(t, 2, 2)
	testDivRoundNearest32(t, 3, 2)
	testDivRoundNearest32(t, 1, 3)
	testDivRoundNearest32(t, 2, 3)
	testDivRoundNearest32(t, 3, 3)
	testDivRoundNearest32(t, 4, 3)

	for i := 0; i < 1000; i++ {
		a := rand.Uint32()
		b := rand.Uint32()
		if b == 0 {
			continue
		}
		testDivRoundNearest32(t, a, b)
	}
}

func min32Ref(x, y int32) int32 {
	if x < y {
		return x
	} else {
		return y
	}
}

func testMin32(t *testing.T, x, y int32) {
	out := Min32(x, y)
	check := min32Ref(x, y)
	if out != check {
		t.Errorf("Min32(%v, %v) != %v (is %v)", x, y, check, out)
	}
}

func TestMin32(t *testing.T) {
	testMin32(t, 1, 1)
	testMin32(t, 1, 2)
	testMin32(t, -1, -0x80000000)
	testMin32(t, -0x7fffffff, 2)

	for i := 0; i < 1000; i++ {
		testMin32(t, int32(rand.Uint32()), int32(rand.Uint32()))
	}
}

func max32Ref(x, y int32) int32 {
	if x > y {
		return x
	} else {
		return y
	}
}

func testMax32(t *testing.T, x, y int32) {
	out := Max32(x, y)
	check := max32Ref(x, y)
	if out != check {
		t.Errorf("Max32(%v, %v) != %v (is %v)", x, y, check, out)
	}
}

func TestMax32(t *testing.T) {
	testMax32(t, 1, 1)
	testMax32(t, 1, 2)
	testMax32(t, -1, -0x80000000)
	testMax32(t, -0x7fffffff, 2)

	for i := 0; i < 1000; i++ {
		testMax32(t, int32(rand.Uint32()), int32(rand.Uint32()))
	}
}

func isPow2Ref32(x uint32) bool {
	if x == 0 {
		return true
	}
	place := uint(0)
	for ; place < 32; place++ {
		if 1<<place == x {
			return true
		}
	}
	return false
}

func testIsPow232(t *testing.T, x uint32) {
	out := IsPow2(uint(x))
	check := isPow2Ref32(x)
	if out != check {
		t.Errorf("IsPow2(0x%x) != %v (is %v)", x, check, out)
	}
}

func isPow2Ref64(x uint64) bool {
	if x == 0 {
		return true
	}
	place := uint(0)
	for ; place < 64; place++ {
		if 1<<place == x {
			return true
		}
	}
	return false
}

func testIsPow264(t *testing.T, x uint64) {
	out := IsPow2(uint(x))
	check := isPow2Ref64(x)
	if out != check {
		t.Errorf("IsPow2(0x%x) != %v (is %v)", x, check, out)
	}
}

func TestIsPow2(t *testing.T) {
	testIsPow232(t, 0)
	testIsPow232(t, 1)
	testIsPow232(t, 2)
	testIsPow232(t, 3)
	testIsPow232(t, 4)
	testIsPow232(t, 0x80000000)
	testIsPow232(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testIsPow232(t, rand.Uint32())
	}

	testIsPow264(t, 0)
	testIsPow264(t, 1)
	testIsPow264(t, 2)
	testIsPow264(t, 3)
	testIsPow264(t, 4)
	testIsPow264(t, 0x80000000)
	testIsPow264(t, 0xffffffff)

	testIsPow264(t, 0x8000000000000000)
	testIsPow264(t, 0xffffffffffffffff)

	for i := 0; i < 1000; i++ {
		testIsPow264(t, fixrand.Uint64())
	}
}

func nlpo232Ref(x uint32) uint32 {
	high := uint32(1 << (32 - 1))
	if x&high == high {
		return 0
	}
	place := uint(0)
	for ; ; place++ {
		b := uint32(1) << place
		if b > x {
			return b
		}
	}
}

func testNlpo232(t *testing.T, x uint32) {
	out := Nlpo232(x)
	check := nlpo232Ref(x)
	if out != check {
		t.Errorf("Nlpo232(%v) != %v (is %v)", x, check, out)
	}
}

func TestNlpo232(t *testing.T) {
	testNlpo232(t, 0)
	testNlpo232(t, 1)
	testNlpo232(t, 2)
	testNlpo232(t, 3)
	testNlpo232(t, 4)
	testNlpo232(t, 5)
	testNlpo232(t, 6)
	testNlpo232(t, 7)
	testNlpo232(t, 8)
	testNlpo232(t, 0x80000000)
	testNlpo232(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testNlpo232(t, rand.Uint32())
	}
}

func nlpo264Ref(x uint64) uint64 {
	high := uint64(1 << (64 - 1))
	if x&high == high {
		return 0
	}
	place := uint(0)
	for ; ; place++ {
		b := uint64(1) << place
		if b > x {
			return b
		}
	}
}

func testNlpo264(t *testing.T, x uint64) {
	out := Nlpo264(x)
	check := nlpo264Ref(x)
	if out != check {
		t.Errorf("Nlpo264(%v) != %v (is %v)", x, check, out)
	}
}

func TestNlpo264(t *testing.T) {
	testNlpo264(t, 0)
	testNlpo264(t, 1)
	testNlpo264(t, 2)
	testNlpo264(t, 3)
	testNlpo264(t, 4)
	testNlpo264(t, 5)
	testNlpo264(t, 6)
	testNlpo264(t, 7)
	testNlpo264(t, 8)
	testNlpo264(t, 0x80000000)
	testNlpo264(t, 0xffffffff)

	testNlpo264(t, 0x8000000000000000)
	testNlpo264(t, 0xffffffffffffffff)

	for i := 0; i < 1000; i++ {
		testNlpo264(t, fixrand.Uint64())
	}
}

func sameWithinTolerance32Ref(a, b, c int32) bool {
	a64 := int64(a)
	b64 := int64(b)
	c64 := int64(c)

	if a64 > b64 {
		return a64-b64 < c64
	} else {
		return b64-a64 < c64
	}
}

func testSameWithinTolerance32(t *testing.T, a, b, c int32) {
	out := SameWithinTolerance32(a, b, c)
	check := sameWithinTolerance32Ref(a, b, c)
	if out != check {
		t.Errorf("SameWithinTolerance32(%v, %v, %v) != %v (is %v)", a, b, c, check, out)
	}
}

func TestSameWithinTolerance32(t *testing.T) {
	testSameWithinTolerance32(t, 3, 2, 1)
	testSameWithinTolerance32(t, 6, 5, 2)
	testSameWithinTolerance32(t, 0, -0x80000000, 1)

	for i := 0; i < 1000; i++ {
		a := int32(rand.Uint32())
		b := int32(rand.Uint32())
		c := int32(rand.Uint32())
		testSameWithinTolerance32(t, a, b, c)
	}
}

func log2Floor32Ref(x uint32) uint32 {
	if x == 0 {
		return 0
	}
	return uint32(math.Log2(float64(x)))
}

func testLog2Floor32(t *testing.T, x uint32) {
	out := Log2Floor32(x)
	check := log2Floor32Ref(x)
	if out != check {
		t.Errorf("Log2Floor32(0x%x) != %v (is %v)", x, check, out)
	}
}

func TestLog2Floor32(t *testing.T) {
	testLog2Floor32(t, 0)
	testLog2Floor32(t, 1)
	testLog2Floor32(t, 2)
	testLog2Floor32(t, 3)
	testLog2Floor32(t, 4)
	testLog2Floor32(t, 5)
	testLog2Floor32(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testLog2Floor32(t, rand.Uint32())
	}
}

func log2Floor64Ref(x uint64) uint64 {
	if x == 0 {
		return 0
	}
	return 64 - lzc64Ref(x) - 1
}

func testLog2Floor64(t *testing.T, x uint64) {
	out := Log2Floor64(x)
	check := log2Floor64Ref(x)
	if out != check {
		t.Errorf("Log2Floor64(0x%x) != %v (is %v)", x, check, out)
	}
}

func TestLog264Floor(t *testing.T) {
	testLog2Floor64(t, 0)
	testLog2Floor64(t, 1)
	testLog2Floor64(t, 2)
	testLog2Floor64(t, 3)
	testLog2Floor64(t, 4)
	testLog2Floor64(t, 5)
	testLog2Floor64(t, 0xffffffff)

	testLog2Floor64(t, 0xffffffffffffffff)
	testLog2Floor64(t, 0xfffffffffffffffe)
	testLog2Floor64(t, 0xffffffffffffffe)
	testLog2Floor64(t, 0xfffffffffffffe)
	testLog2Floor64(t, 0xffffffffffffe)
	testLog2Floor64(t, 0xfffffffffffe)

	for i := 0; i < 1000; i++ {
		testLog2Floor64(t, fixrand.Uint64())
	}
}

func log2Ceil32Ref(x uint32) uint32 {
	if x == 0 {
		return 0
	}
	return uint32(math.Ceil(math.Log2(float64(x))))
}

func testLog2Ceil32(t *testing.T, x uint32) {
	out := Log2Ceil32(x)
	check := log2Ceil32Ref(x)
	if out != check {
		t.Errorf("Log2Ceil32(0x%x) != %v (is %v)", x, check, out)
	}
}

func TestLog2Ceil32(t *testing.T) {
	testLog2Ceil32(t, 0)
	testLog2Ceil32(t, 1)
	testLog2Ceil32(t, 2)
	testLog2Ceil32(t, 3)
	testLog2Ceil32(t, 4)
	testLog2Ceil32(t, 5)
	testLog2Ceil32(t, 6)
	testLog2Ceil32(t, 7)
	testLog2Ceil32(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testLog2Ceil32(t, rand.Uint32())
	}
}

func log2Ceil64Ref(x uint64) uint64 {
	if x == 0 {
		return 0
	}
	r := 64 - lzc64Ref(x)
	if isPow2Ref64(x) {
		r--
	}
	return r
}

func testLog2Ceil64(t *testing.T, x uint64) {
	out := Log2Ceil64(x)
	check := log2Ceil64Ref(x)
	if out != check {
		t.Errorf("Log2Ceil64(0x%x) != %v (is %v)", x, check, out)
	}
}

func TestLog2Ceil64(t *testing.T) {
	testLog2Ceil64(t, 0)
	testLog2Ceil64(t, 1)
	testLog2Ceil64(t, 2)
	testLog2Ceil64(t, 3)
	testLog2Ceil64(t, 4)
	testLog2Ceil64(t, 5)
	testLog2Ceil64(t, 6)
	testLog2Ceil64(t, 7)
	testLog2Ceil64(t, 8)
	testLog2Ceil64(t, 0xffffffff)

	for i := 0; i < 1000; i++ {
		testLog2Ceil64(t, fixrand.Uint64())
	}
}
