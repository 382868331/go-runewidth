package runewidth

import "testing"

// TestTask011 isolates the 并发宽度统计 regression.
func TestTask011(t *testing.T) {
	TestIsAmbiguousWidth(t)
}

// TestTask011Repeat guards deterministic behavior across repeated calls.
func TestTask011Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask011(t)
	}
}
