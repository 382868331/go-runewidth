package runewidth

import "testing"

// TestTask014 isolates the 终端布局 regression.
func TestTask014(t *testing.T) {
	TestTruncateSmaller(t)
}

// TestTask014Repeat guards deterministic behavior across repeated calls.
func TestTask014Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask014(t)
	}
}
