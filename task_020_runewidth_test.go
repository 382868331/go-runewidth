package runewidth

import "testing"

// TestTask020 isolates the Unicode版本 regression.
func TestTask020(t *testing.T) {
	TestTruncateLeft(t)
}
