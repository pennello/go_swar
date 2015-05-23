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
