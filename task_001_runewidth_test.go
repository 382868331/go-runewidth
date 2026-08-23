package runewidth

import "testing"

// TestTask001 isolates the 空文本宽度 regression.
func TestTask001(t *testing.T) {
	TestIsEastAsian(t)
}
