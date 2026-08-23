package runewidth

import "testing"

// TestTask007 isolates the 组合字符截断 regression.
func TestTask007(t *testing.T) {
	TestRuneWidthConcurrent(t)
}
