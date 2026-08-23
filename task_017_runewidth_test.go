package runewidth

import "testing"

// TestTask017 isolates the emoji边界 regression.
func TestTask017(t *testing.T) {
	TestTruncateJustFit(t)
}

// TestTask017Repeat guards deterministic behavior across repeated calls.
func TestTask017Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask017(t)
	}
}
