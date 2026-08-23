package runewidth

import "testing"

// TestTask013 isolates the Windows换行 regression.
func TestTask013(t *testing.T) {
	TestStringWidthInvalid(t)
}

// TestTask013Repeat guards deterministic behavior across repeated calls.
func TestTask013Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask013(t)
	}
}
