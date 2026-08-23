package runewidth

import "testing"

// TestTask015 isolates the 终端句柄 regression.
func TestTask015(t *testing.T) {
	TestTruncate(t)
}
