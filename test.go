// chris 052315 Common code for testing.

package swar

// Frequently, a complaint about testing is that you're implementing
// your logic twice.  In the case of these swar algorithms, though,
// that's exactly what we want to do!  These algorithms implement some
// more trivial logic in a weird, interesting, or efficient way.
// Therefore, our approach is to implement reference functions that
// implement the same logic in a more normal way, and then just check
// the two against each other.

import (
	"os"
	"testing"
	"time"

	"math/rand"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
	os.Exit(m.Run())
}
