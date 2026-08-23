package runewidth

import "testing"

// TestTask011 isolates the 并发宽度统计 regression.
func TestTask011(t *testing.T) {
	TestIsAmbiguousWidth(t)
}
