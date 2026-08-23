package runewidth

import "testing"

// TestTask003 isolates the 宽度表解析 regression.
func TestTask003(t *testing.T) {
	TestIsEastAsianLANG(t)
}

// TestTask003Repeat guards deterministic behavior across repeated calls.
func TestTask003Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask003(t)
	}
}
