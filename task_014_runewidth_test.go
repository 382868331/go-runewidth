package runewidth

import "testing"

// TestTask014 isolates the 终端布局 regression.
func TestTask014(t *testing.T) {
	TestTruncateSmaller(t)
}
