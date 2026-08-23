package runewidth

import "testing"

// TestTask016 isolates the 宽度模式恢复 regression.
func TestTask016(t *testing.T) {
	TestTruncateFit(t)
}

// TestTask016Repeat guards deterministic behavior across repeated calls.
func TestTask016Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask016(t)
	}
}
