// chris 052315

package swar

import (
	"testing"

	"math/rand"
)

func integerSelect32Ref(a, b, c, d int32) int32 {
	if a < b {
		return c
	} else {
		return d
	}
}

func testIntegerSelect32(t *testing.T, a, b, c, d int32) {
	out := IntegerSelect32(a, b, c, d)
	check := integerSelect32Ref(a, b, c, d)
	if out != check {
		t.Errorf("IntegerSelect32(%v, %v, %v, %v) != %v (is %v)", a, b, c, d, check, out)
	}
}

func TestIntegerSelect32(t *testing.T) {
	testIntegerSelect32(t, 1, 2, 3, 4)
	testIntegerSelect32(t, 2, 1, 3, 4)
	testIntegerSelect32(t, -0x80000000, 1, 3, 4)
	testIntegerSelect32(t, 0x7fffffff, -1, 3, 4)

	for i := 0; i < 1000; i++ {
		a := int32(rand.Uint32())
		b := int32(rand.Uint32())
		c := int32(rand.Uint32())
		d := int32(rand.Uint32())
		testIntegerSelect32(t, a, b, c, d)
	}
}
