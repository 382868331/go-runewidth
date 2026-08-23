package runewidth

import "testing"

// TestTask006 isolates the 字符类别去重 regression.
func TestTask006(t *testing.T) {
	TestStrictWidthLUT(t)
}
