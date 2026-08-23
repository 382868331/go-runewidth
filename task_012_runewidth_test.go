package runewidth

import "testing"

// TestTask012 isolates the ANSI文本宽度 regression.
func TestTask012(t *testing.T) {
	TestStringWidth(t)
}
