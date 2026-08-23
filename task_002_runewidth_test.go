package runewidth

import "testing"

// TestTask002 isolates the rune切片 regression.
func TestTask002(t *testing.T) {
	TestIsEastAsianLCCTYPE(t)
}

// TestTask002Repeat guards deterministic behavior across repeated calls.
func TestTask002Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask002(t)
	}
}
