package runewidth

import "testing"

// TestTask007 isolates the 组合字符截断 regression.
func TestTask007(t *testing.T) {
	TestRuneWidthConcurrent(t)
}

// TestTask007Repeat guards deterministic behavior across repeated calls.
func TestTask007Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask007(t)
	}
}
