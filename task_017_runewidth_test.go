package runewidth

import "testing"

// TestTask017 isolates the emoji边界 regression.
func TestTask017(t *testing.T) {
	TestTruncateJustFit(t)
}
