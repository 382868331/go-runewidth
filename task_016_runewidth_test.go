package runewidth

import "testing"

// TestTask016 isolates the 宽度模式恢复 regression.
func TestTask016(t *testing.T) {
	TestTruncateFit(t)
}
