package runewidth

import "testing"

// TestTask012 isolates the ANSI文本宽度 regression.
func TestTask012(t *testing.T) {
	TestStringWidth(t)
}

// TestTask012Repeat guards deterministic behavior across repeated calls.
func TestTask012Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask012(t)
	}
}
