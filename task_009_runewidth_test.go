package runewidth

import "testing"

// TestTask009 isolates the 宽高字段 regression.
func TestTask009(t *testing.T) {
	TestSorted(t)
}

// TestTask009Repeat guards deterministic behavior across repeated calls.
func TestTask009Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask009(t)
	}
}
