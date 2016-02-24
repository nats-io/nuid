// +build !race

package nuid

import (
	"testing"
)

func TestBasicUniqueness(t *testing.T) {
	n := 10000000
	m := make(map[string]bool)
	for i := 0; i < n; i++ {
		n := Next()
		if m[n] {
			t.Fatalf("Duplicate NUID found: %v\n", n)
		}
	}
}
