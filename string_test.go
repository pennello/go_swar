// chris 052315

package swar

import (
	"strings"
	"testing"
)

func testAsciiUpper(t *testing.T, s string) {
	out := AsciiUpper(s)
	check := strings.ToUpper(s)
	if out != check {
		t.Errorf("AsciiUpper(%v) != %v (is %v)", s, check, out)
	}
}

func TestAsciiUpper(t *testing.T) {
	testAsciiUpper(t, "")
	testAsciiUpper(t, "abcdefghijklmnopqrstuvwxyz")
	testAsciiUpper(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func testAsciiLower(t *testing.T, s string) {
	out := AsciiLower(s)
	check := strings.ToLower(s)
	if out != check {
		t.Errorf("AsciiLower(%v) != %v (is %v)", s, check, out)
	}
}

func TestAsciiLower(t *testing.T) {
	testAsciiLower(t, "")
	testAsciiLower(t, "abcdefghijklmnopqrstuvwxyz")
	testAsciiLower(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
