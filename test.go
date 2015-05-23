// chris 052315 Common code for testing.

package swar

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
